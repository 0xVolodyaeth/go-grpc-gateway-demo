package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"proxy/service_proto"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	// service_proto.RegisterYourServiceHandlerServer(ctx, mux, &service_proto.YourServiceHandler{})
	err := service_proto.RegisterYourServiceHandlerFromEndpoint(ctx, mux, "localhost:3000", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("gateway server has started on port 8081")
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
