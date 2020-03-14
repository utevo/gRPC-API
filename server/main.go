package main

import (
	context "context"
	net "net"

	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"

	proto "github.com/utevo/gRPC-API/proto"
)

type server struct{}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}


func main() {
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}