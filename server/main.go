package main

import (
	"context"
	"gogrpcgateway/service_proto"
	"log"
	"net"

	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	// _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	"google.golang.org/grpc"
	// _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	// _ "google.golang.org/protobuf/cmd/protoc-gen-go
)

type ExampleService struct {
	*service_proto.UnimplementedYourServiceServer
}

func main() {
	exs := ExampleService{}

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()

	service_proto.RegisterYourServiceServer(srv, &exs)

	log.Println("GRPC is listening on port 3000")
	if err = srv.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func (ex *ExampleService) Echo(ctx context.Context, msg *service_proto.StringMessage) (*service_proto.StringMessage, error) {

	return &service_proto.StringMessage{
		Value: "Hello " + msg.Value,
	}, nil
}
