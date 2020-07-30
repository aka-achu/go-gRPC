package client

import (
	"context"
	"fmt"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"os"
	"path/filepath"
	"time"
)

func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var credentialConfigStatus bool
	//Checking for existence of PerRPCCredsCallOption
	for _, o := range opts {
		_, ok := o.(grpc.PerRPCCredsCallOption)
		if ok {
			credentialConfigStatus = true
			break
		}
	}
	// If PerRPCCredsCallOption is not present
	if !credentialConfigStatus {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: os.Getenv("FALLBACK_TOKEN"),
		})))
	}
	logger("Unary RPC Request: %s, Time-%v", method, time.Now())
	return invoker(ctx, method, req, reply, cc, opts...)
}

func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credentialConfigStatus bool
	//Checking for existence of PerRPCCredsCallOption
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			credentialConfigStatus = true
			break
		}
	}
	// If PerRPCCredsCallOption is not present
	if !credentialConfigStatus {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: os.Getenv("FALLBACK_TOKEN"),
		})))
	}
	logger("Stream RPC Request: %s, Time-%v", method, time.Now())
	if clientStream, err := streamer(ctx, desc, cc, method, opts...); err != nil {
		return nil, err
	} else {
		return &struct {
			grpc.ClientStream
		}{clientStream,
		}, nil
	}
}

func Initialize() {
	r := manual.NewBuilderWithScheme("operation")
	r.InitialState(resolver.State{
		Addresses: []resolver.Address{
			{Addr: os.Getenv("SERVER_1_ADDRESS")},
			{Addr: os.Getenv("SERVER_2_ADDRESS")},
		},
	})

	opts := grpc.WithInsecure()
	if os.Getenv("SSL_MODE") == "true" {
		if credential, err := credentials.NewClientTLSFromFile(
			filepath.Join(
				os.Getenv("CERTIFICATE_DIR"),
				os.Getenv("CA_FILE"),
			),
			"localhost",
		); err != nil {
			fatalLogger("Failed to load CA trust certificate. -%v", err)
			return
		} else {
			opts = grpc.WithTransportCredentials(credential)
		}

	}
	if clientConnection, err := grpc.Dial(
		fmt.Sprintf("%s:///unused", r.Scheme()),
		opts,
		grpc.WithBlock(),
		grpc.WithUnaryInterceptor(unaryInterceptor),
		grpc.WithStreamInterceptor(streamInterceptor),
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		); err != nil {
		fatalLogger("could not connect: %v", err)
	} else {
		defer clientConnection.Close()
		c := operation_pb.NewOperationServiceClient(clientConnection)
		Sum(c)
		PrimeFactors(c)
		ComputeAverage(c)
		FloorCeiling(c)
		//SquareRoot(c)
		//Power(c, 1*time.Second)
		Power(c, 5*time.Second)
		SumWithCompressor(c)
	}
}
