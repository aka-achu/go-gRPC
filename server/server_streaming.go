package server

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"log"
)

func (*UnimplementedOperationService) PrimeFactors(
	request *operation_pb.PrimeFactorsRequest,
	stream operation_pb.OperationService_PrimeFactorsServer,
) error {
	var number = request.GetNumber()
	for factor := int64(2); number > 1; factor++ {
		if number%factor == 0 {
			if err := stream.Send(&operation_pb.PrimeFactorsResponse{Number: factor}); err != nil {
				log.Fatalf("Failed to stream the factor to the client. - %v", err)
			}
			number = number / factor
		}
	}
	return nil
}
