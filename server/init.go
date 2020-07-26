package server

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UnimplementedOperationService struct{}

func Initialize() {
	if listener, err := net.Listen("tcp", "0.0.0.0:50051"); err != nil {
		log.Fatalf("Failed to listen @'0.0.0.0:50051'-  %v", err)
	} else {
		s := grpc.NewServer()
		operation_pb.RegisterOperationServiceServer(s, &UnimplementedOperationService{})
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed start the serve-r %v", err)
		}
	}
}
