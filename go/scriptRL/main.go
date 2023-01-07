package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math"
	"time"

	colpb "github.com/scionproto/scion/go/pkg/proto/colibri"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = 5045

func main() {

	var serverAddr string
	var address string
	var cbs int64
	var rate float64
	var ingress int
	var egress int

	flag.StringVar(&serverAddr, "s", "", "address of the server")
	flag.StringVar(&address, "address", "", "address of an AS")
	flag.Int64Var(&cbs, "cbs", -1, "cbs")
	flag.Float64Var(&rate, "rate", -1, "rate")
	flag.IntVar(&ingress, "ingress", -1, "ingress")
	flag.IntVar(&egress, "egress", -1, "egress")
	flag.Parse()

	if address == "" {
		log.Fatalln(fmt.Errorf("identifer empty"))
	}

	if ingress < 0 || ingress > math.MaxUint16 {
		log.Fatalln(fmt.Errorf(fmt.Sprintf("ingress needs to be between 0 and %d", math.MaxUint16)))
	}

	if egress < 0 || egress > math.MaxUint16 {
		log.Fatalln(fmt.Errorf(fmt.Sprintf("egress needs to be between 0 and %d", math.MaxUint16)))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", serverAddr, port), opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := colpb.NewRateLimiterServiceClient(conn)

	if cbs != -1 && rate != -1 {
		_, err = client.SetBurstSizeAndRate(ctx, &colpb.SetBurstSizeAndRateRequest{Address: address,
			Cbs: cbs, Rate: rate, Ingress: int32(ingress), Egress: int32(egress)})
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	if cbs != -1 {
		_, err = client.SetBurstSize(ctx, &colpb.SetBurstSizeRequest{Address: address, Cbs: cbs,
			Ingress: int32(ingress), Egress: int32(egress)})
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	if rate != -1 {
		_, err = client.SetRate(ctx, &colpb.SetRateRequest{Address: address, Rate: rate,
			Ingress: int32(ingress), Egress: int32(egress)})
		if err != nil {
			log.Fatalln(err)
		}
	}

}
