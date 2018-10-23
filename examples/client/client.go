package main

import (
	"context"
	"fmt"
	"github.com/nosixtools/LearnGrpc/discovery/resolver"
	"github.com/nosixtools/LearnGrpc/examples/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"time"
)

func main() {

	schema, err := resolver.GenerateAndRegisterConsulResolver("127.0.0.1:8500", "HelloService")
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:///HelloService", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewHelloServiceClient(conn)

	// Contact the server and print out its response.
	name := "nosixtools"

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		r, err := c.SayHello(ctx, &proto.HelloRequest{Name: name})
		if err != nil {
			log.Println("could not greet: %v", err)

		} else {
			log.Printf("Hello: %s", r.Result)
		}
		time.Sleep(time.Second)
	}

}
