package client

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Initialize() {
	opts := grpc.WithInsecure()
	if os.Getenv("SSL_MODE") == "true" {
		if credential, err := credentials.NewClientTLSFromFile(
			filepath.Join(
				os.Getenv("CERTIFICATE_DIR"),
				os.Getenv("CA_FILE"),
			),
			"localhost",
		); err != nil {
			log.Fatalf("Failed to load CA trust certificate. -%v", err)
			return
		} else {
			opts = grpc.WithTransportCredentials(credential)
		}

	}
	if clientConnection, err := grpc.Dial(os.Getenv("SERVER_ADDRESS"), opts); err != nil {
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
