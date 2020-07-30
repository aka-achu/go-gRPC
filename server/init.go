package server

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
	"path/filepath"
)

type Service struct {
	operation_pb.UnimplementedOperationServiceServer
}

func Initialize() {

	if listener, err := net.Listen("tcp", os.Getenv("SERVER_ADDRESS")); err != nil {
		log.Fatalf("Failed to listen @'%s'-  %v", os.Getenv("SERVER_ADDRESS"), err)
	} else {
		var opts []grpc.ServerOption
		if os.Getenv("SSL_MODE") == "true" {
			if credential, err := credentials.NewServerTLSFromFile(
				filepath.Join(
					os.Getenv("CERTIFICATE_DIR"),
					os.Getenv("CERTIFICATE_FILE"),
				), filepath.Join(
					os.Getenv("CERTIFICATE_DIR"),
					os.Getenv("KEY_FILE"),
				),
			); err != nil {
				log.Fatalf("Failed to load ssl certificates. -%v", err)
			} else {
				opts = append(opts, grpc.Creds(credential))
			}
		}
		s := grpc.NewServer(opts...)
		operation_pb.RegisterOperationServiceServer(s, &Service{})
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed start the serve-r %v", err)
		}
	}
}
