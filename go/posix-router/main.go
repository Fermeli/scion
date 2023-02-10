// Copyright 2020 Anapaya Systems
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/scionproto/scion/go/lib/log"
	"github.com/scionproto/scion/go/lib/serrors"
	"github.com/scionproto/scion/go/lib/topology"
	"github.com/scionproto/scion/go/pkg/app"
	"github.com/scionproto/scion/go/pkg/app/launcher"
	colpb "github.com/scionproto/scion/go/pkg/proto/colibri"
	"github.com/scionproto/scion/go/pkg/router"
	"github.com/scionproto/scion/go/pkg/router/api"
	"github.com/scionproto/scion/go/pkg/router/config"
	"github.com/scionproto/scion/go/pkg/router/control"
	"github.com/scionproto/scion/go/pkg/service"
)

const port = 5045

var globalCfg config.Config

type rateLimiterServer struct {
	colpb.UnimplementedRateLimiterServiceServer
	dp *router.DataPlane
}

func (r *rateLimiterServer) SetBurstSize(ctx context.Context, req *colpb.SetBurstSizeRequest) (*colpb.Success, error) {
	rateLimiter, ok := r.dp.SyncRateLimiters[uint16(req.Ingress)]

	if !ok {
		r.dp.InitRateLimiter(uint16(req.Ingress))
		rateLimiter, _ = r.dp.SyncRateLimiters[uint16(req.Ingress)]
	}

	rateLimiter.Lock()
	defer rateLimiter.Unlock()
	identifier, err := r.dp.BuildIdentifier(uint16(req.Egress), req.Address)

	if err != nil {
		return &colpb.Success{}, err
	}
	err = rateLimiter.Ratelimiter.SetBurstSize(identifier, req.Cbs)

	if err != nil {
		return &colpb.Success{}, err
	}

	err = r.dp.SetCbsMetric(uint16(req.Ingress), uint16(req.Egress), req.Address, uint64(req.Cbs))

	return &colpb.Success{}, err
}

func (r *rateLimiterServer) SetBurstSizeAndRate(ctx context.Context, req *colpb.SetBurstSizeAndRateRequest) (*colpb.Success, error) {
	rateLimiter, ok := r.dp.SyncRateLimiters[uint16(req.Ingress)]

	if !ok {
		r.dp.InitRateLimiter(uint16(req.Ingress))
		rateLimiter, _ = r.dp.SyncRateLimiters[uint16(req.Ingress)]
		return &colpb.Success{}, fmt.Errorf("aaa")
	}

	rateLimiter.Lock()
	defer rateLimiter.Unlock()
	identifier, err := r.dp.BuildIdentifier(uint16(req.Egress), req.Address)

	if err != nil {
		return &colpb.Success{}, err
	}

	if !rateLimiter.Ratelimiter.Contains(identifier) {
		err = r.dp.InitRateLimiterMetrics(uint16(req.Ingress), uint16(req.Egress), req.Address, uint64(req.Cbs), req.Rate)
		if err != nil {
			return &colpb.Success{}, err
		}
		rateLimiter.Ratelimiter.AddRatelimit(identifier, req.Rate, req.Cbs, time.Now())
		return &colpb.Success{}, nil
	}
	err = rateLimiter.Ratelimiter.SetBurstSizeAndRate(identifier, req.Cbs, req.Rate)
	if err != nil {
		return &colpb.Success{}, err
	}
	err = r.dp.SetCbsMetric(uint16(req.Ingress), uint16(req.Egress), req.Address, uint64(req.Cbs))
	if err != nil {
		return &colpb.Success{}, err
	}
	r.dp.SetRateMetric(uint16(req.Ingress), uint16(req.Egress), req.Address, req.Rate)

	return &colpb.Success{}, err
}

func (r *rateLimiterServer) SetRate(ctx context.Context, req *colpb.SetRateRequest) (*colpb.Success, error) {
	rateLimiter, ok := r.dp.SyncRateLimiters[uint16(req.Ingress)]

	if !ok {
		r.dp.InitRateLimiter(uint16(req.Ingress))
		rateLimiter, _ = r.dp.SyncRateLimiters[uint16(req.Ingress)]
	}

	rateLimiter.Lock()
	defer rateLimiter.Unlock()
	identifier, err := r.dp.BuildIdentifier(uint16(req.Egress), req.Address)

	if err != nil {
		return &colpb.Success{}, err
	}
	err = rateLimiter.Ratelimiter.SetRate(identifier, req.Rate)
	if err != nil {
		return &colpb.Success{}, err
	}
	err = r.dp.SetRateMetric(uint16(req.Ingress), uint16(req.Egress), req.Address, req.Rate)

	return &colpb.Success{}, err
}

func main() {
	application := launcher.Application{
		TOMLConfig: &globalCfg,
		ShortName:  "SCION Router",
		Main:       realMain,
	}
	application.Run()
}

func realMain(ctx context.Context) error {
	controlConfig, err := loadControlConfig()
	if err != nil {
		return err
	}
	g, errCtx := errgroup.WithContext(ctx)
	metrics := router.NewMetrics()
	dp := &router.Connector{
		DataPlane: router.DataPlane{
			Metrics: metrics,
		},
	}
	iaCtx := &control.IACtx{
		Config: controlConfig,
		DP:     dp,
	}
	if err := iaCtx.Configure(); err != nil {
		return serrors.WrapStr("configuring dataplane", err)
	}
	statusPages := service.StatusPages{
		"info":      service.NewInfoStatusPage(),
		"config":    service.NewConfigStatusPage(globalCfg),
		"log/level": service.NewLogLevelStatusPage(),
		"topology":  topologyHandler(iaCtx.Config.Topo),
	}
	if err := statusPages.Register(http.DefaultServeMux, globalCfg.General.ID); err != nil {
		return err
	}

	var cleanup app.Cleanup
	g.Go(func() error {
		defer log.HandlePanic()
		<-errCtx.Done()
		return cleanup.Do()
	})

	// Initialize and start service management API.
	if globalCfg.API.Addr != "" {
		r := chi.NewRouter()
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
		}))
		r.Get("/", api.ServeSpecInteractive)
		r.Get("/openapi.json", api.ServeSpecJSON)
		server := api.Server{
			Config:    service.NewConfigStatusPage(globalCfg).Handler,
			Info:      service.NewInfoStatusPage().Handler,
			LogLevel:  service.NewLogLevelStatusPage().Handler,
			Dataplane: dp,
		}
		log.Info("Exposing API", "addr", globalCfg.API.Addr)
		h := api.HandlerFromMuxWithBaseURL(&server, r, "/api/v1")
		mgmtServer := &http.Server{
			Addr:    globalCfg.API.Addr,
			Handler: h,
		}
		cleanup.Add(mgmtServer.Close)
		g.Go(func() error {
			defer log.HandlePanic()
			err := mgmtServer.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return serrors.WrapStr("serving service management API", err)
			}
			return nil
		})
	}
	g.Go(func() error {
		defer log.HandlePanic()
		return globalCfg.Metrics.ServePrometheus(errCtx)
	})
	g.Go(func() error {
		defer log.HandlePanic()
		if err := dp.DataPlane.Run(errCtx); err != nil {
			return serrors.WrapStr("running dataplane", err)
		}
		return nil
	})

	// Run a grpc server to listen to rate limit adjsutment requests
	g.Go(func() error {
		defer log.HandlePanic()
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", controlConfig.BR.InternalAddr.IP.String(), port)) //grpcServerAddr)
		if err != nil {
			return serrors.WrapStr("failed to listen:", err)
		}
		var opts []grpc.ServerOption
		rateLimiterServer := rateLimiterServer{dp: &dp.DataPlane}
		grpcServer := grpc.NewServer(opts...)
		colpb.RegisterRateLimiterServiceServer(grpcServer, &rateLimiterServer)
		grpcServer.Serve(lis)
		return nil
	})

	return g.Wait()
}

func loadControlConfig() (*control.Config, error) {
	newConf, err := control.LoadConfig(globalCfg.General.ID, globalCfg.General.ConfigDir)
	if err != nil {
		return nil, serrors.WrapStr("loading topology", err)
	}
	return newConf, nil
}

func topologyHandler(topo topology.Topology) service.StatusPage {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		bytes, err := json.MarshalIndent(topo, "", "    ")
		if err != nil {
			http.Error(w, "Unable to marshal topology", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(bytes)+"\n")
	}
	return service.StatusPage{
		Info:    "SCION topology",
		Handler: handler,
	}
}
