package client

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func Initialize() {
	if clientConnection, err := grpc.Dial(os.Getenv("SERVER_ADDRESS"), grpc.WithInsecure()); err != nil {
		log.Fatalf("could not connect: %v", err)
	} else {
		defer clientConnection.Close()
		c := operation_pb.NewOperationServiceClient(clientConnection)
		Sum(c)
		PrimeFactors(c)
		ComputeAverage(c)
		FloorCeiling(c)
		SquareRoot(c)
		Power(c, 1*time.Second)
		Power(c, 5*time.Second)
	}
}
