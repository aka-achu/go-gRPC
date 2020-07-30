package server

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Service struct {
	operation_pb.UnimplementedOperationServiceServer
}

func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	return token == os.Getenv("FALLBACK_TOKEN")
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return m, err
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
		opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor))
		s := grpc.NewServer(opts...)
		operation_pb.RegisterOperationServiceServer(s, &Service{})
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed start the serve-r %v", err)
		}
	}
}
