package server

import (
	"context"
	"pb/gen/pb"
)

type GreetServer struct {
	greetServiceV1.UnimplementedGreetServiceServer
}

func (g GreetServer) Greet(_ context.Context, request *greetServiceV1.GreetRequest) (*greetServiceV1.GreetResponse, error) {
	return &greetServiceV1.GreetResponse{
		Message: "8080 " + request.Message,
	}, nil
}
