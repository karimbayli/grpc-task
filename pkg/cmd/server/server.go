package cmd

import (
	"context"
	"flag"
	"fmt"
	"mgidAssignment/pkg/protocol/grpc"
	"mgidAssignment/pkg/service/v1"
)

type Config struct {
	// gRPC server start parameters section
	// gRPC is TCP port to listen by gRPC server
	GRPCPort string
}

func RunServer() error {
	ctx := context.Background()
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.Parse()
	v1API := v1.NewService()
	fmt.Println(cfg)

	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}
