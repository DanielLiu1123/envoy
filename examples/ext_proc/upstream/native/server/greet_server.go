package server

import (
	"context"
	"pb/gen/pb"
)

type GreetServer struct {
	greet_service_v1.UnimplementedGreetServiceServer
}

func (g GreetServer) Greet(ctx context.Context, request *greet_service_v1.GreetRequest) (*greet_service_v1.GreetResponse, error) {
	return &greet_service_v1.GreetResponse{
		Message: "8080 " + request.Message,
	}, nil
}
