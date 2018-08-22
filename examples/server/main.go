package main

import (
	"fmt"
	"net"
	"github.com/nosixtools/LearnGrpc/examples/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/nosixtools/LearnGrpc/discovery"
	"time"
	"github.com/nosixtools/LearnGrpc/discovery/register"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("client called! 8081")
	return &proto.HelloResponse{Result: "hi," + in.Name + "!"}, nil
}

const (
	host        = "127.0.0.1"
	port        = 8081
	consul_port = 8500
)

func main() {

	listen, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(host), port, "",})
	if err != nil {
		fmt.Println(err.Error())
	}
	s := grpc.NewServer()

	// register service
	cr := register.NewConsulRegister(fmt.Sprintf("%s:%d", host, consul_port), 15)
	cr.Register(discovery.RegisterInfo{
		Host:           host,
		Port:           port,
		ServiceName:    "HelloService",
		UpdateInterval: time.Second})

	proto.RegisterHelloServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		fmt.Println("failed to serve:" + err.Error())
	}
}
