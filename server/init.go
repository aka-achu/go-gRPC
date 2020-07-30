package server

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
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
		fatalLogger("RPC failed with error. -%v", err)
	}
	logger("Unary RPC Request: Time-%v", time.Now())
	return m, err
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}
	if err := handler(srv, &struct {
		grpc.ServerStream
	}{
		ss,
	});err != nil {
		fatalLogger("RPC failed with error. -%v", err)
		return err
	} else {
		logger("Stream RPC Request: Time-%v", time.Now())
		return nil
	}
}

func Serve(serverAddress string) {
	if listener, err := net.Listen("tcp", serverAddress); err != nil {
		fatalLogger("Failed to listen @'%s'-  %v", serverAddress, err)
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
				fatalLogger("Failed to load ssl certificates. -%v", err)
			} else {
				opts = append(opts, grpc.Creds(credential))
			}
		}
		opts = append(opts, grpc.UnaryInterceptor(unaryInterceptor))
		opts = append(opts, grpc.StreamInterceptor(streamInterceptor))
		s := grpc.NewServer(opts...)
		operation_pb.RegisterOperationServiceServer(s, &Service{})
		if err := s.Serve(listener); err != nil {
			fatalLogger("failed start the serve-r %v", err)
		}
	}
}

func Initialize() {
	var wg sync.WaitGroup
	for _, addr := range []string{os.Getenv("SERVER_1_ADDRESS"), os.Getenv("SERVER_2_ADDRESS")} {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			Serve(addr)
		}(addr)
	}
	wg.Wait()
}
