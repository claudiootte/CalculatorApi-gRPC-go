package main

import (
	"context"
	"fmt"
	"net"

	"github.com/claudiootte/CalculatorApi-gRPC-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	fmt.Println("Connecting to server...")
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Sum(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum01(), request.GetNum02()

	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum01(), request.GetNum02()

	result := a - b

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum01(), request.GetNum02()

	result := a * b

	return &proto.Response{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetNum01(), request.GetNum02()

	result := a / b

	return &proto.Response{Result: result}, nil

}
