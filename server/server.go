package main

import (
	"context"
	"fmt"
	"net"

	serviceServer "github.com/maxpattmanqa/proto/go/serverservice2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	serviceServer.UnimplementedAddServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err.Error)
	}
	srv := grpc.NewServer()
	//serviceServer.mustEmbedUnimplementedAddServiceServer()

	serviceServer.RegisterAddServiceServer(srv, &server{})
	// proto.RegisterAddService(srv, &Server{})

	reflection.Register(srv)
	fmt.Println("things seem to be working ")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *serviceServer.Request) (*serviceServer.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &serviceServer.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *serviceServer.Request) (*serviceServer.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &serviceServer.Response{Result: result}, nil
}
