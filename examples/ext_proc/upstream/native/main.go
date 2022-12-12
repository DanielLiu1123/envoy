package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	greetServiceV1 "pb/gen/pb"
	server2 "pb/native/server"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	greetServiceV1.RegisterGreetServiceServer(server, &server2.GreetServer{})

	reflection.Register(server)

	fmt.Println("server started at " + listen.Addr().String())
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
