package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	colpb "github.com/scionproto/scion/go/pkg/proto/colibri"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var identifier string
	var cbs int64
	var rate float64
	var add bool

	flag.StringVar(&identifier, "identifier", "", "address of an AS")
	flag.Int64Var(&cbs, "cbs", -1, "cbs")
	flag.Float64Var(&rate, "rate", -1, "rate")
	flag.BoolVar(&add, "add", false, "add")
	flag.Parse()

	if identifier == "" {
		log.Fatalln(fmt.Errorf("identifer empty"))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("step 1")
	conn, err := grpc.Dial("localhost:5045", opts...)
	fmt.Println("step 2")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := colpb.NewRateLimiterServiceClient(conn)

	if add {
		if cbs == -1 || rate == -1 {
			log.Fatalln("The cbs and the rate must be set to add a new rate limit")
		}
		_, err = client.AddRateLimit(ctx, &colpb.AddRateLimitRequest{Identifier: identifier, Cbs: cbs, Rate: rate})

		if err != nil {
			log.Fatalln(err)
		}

		return
	}

	if cbs != -1 && rate != -1 {
		_, err = client.SetBurstSizeAndRate(ctx, &colpb.SetBurstSizeAndRateRequest{Identifier: identifier, Cbs: cbs, Rate: rate})
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	if cbs != -1 {
		_, err = client.SetBurstSize(ctx, &colpb.SetBurstSizeRequest{Identifier: identifier, Cbs: cbs})
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	if rate != -1 {
		_, err = client.SetRate(ctx, &colpb.SetRateRequest{Identifier: identifier, Rate: rate})
		if err != nil {
			log.Fatalln(err)
		}
	}

}
