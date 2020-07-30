package server

import (
	"context"
	"fmt"
	"github.com/aka-achu/go-gRPC/models/greet_pb"
)

func (*GreetService) Greet(
	ctx context.Context,
	request *greet_pb.GreetRequest,
) (
	*greet_pb.GreetResponse,
	error,
) {
	return &greet_pb.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s", request.GetName()),
		},nil
}