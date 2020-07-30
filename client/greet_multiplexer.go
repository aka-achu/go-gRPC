package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/greet_pb"
)

func Greet(c greet_pb.GreetServiceClient) {
	if response, err := c.Greet(
		context.Background(),
		&greet_pb.GreetRequest{
			Name: "Achyuta Das",
		}); err != nil {
		fatalLogger("Failed to make the Sum RPC request- %v", err)
	} else {
		printer("Response from server- Greeting %s", response.GetGreeting())
	}
}