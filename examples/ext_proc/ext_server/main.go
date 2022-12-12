package main

import (
	"fmt"
	ext_procv3 "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3"
	"google.golang.org/grpc"
	server2 "moego-grey-gateway/server"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9999))
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	externalProcessingServer := server2.NewExternalProcessingServer()
	ext_procv3.RegisterExternalProcessorServer(server, externalProcessingServer)

	fmt.Println("server started at " + lis.Addr().String())
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
