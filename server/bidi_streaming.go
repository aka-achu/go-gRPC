package server

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"io"
	"log"
	"math"
)

func (*UnimplementedOperationService) FloorCeiling(stream operation_pb.OperationService_FloorCeilingServer) error {
	for {
		if request, err := stream.Recv(); err == io.EOF {
			return nil
		} else if err != nil {
			log.Fatalf("Failed to read client stream. -%v", err)
			return err
		} else {
			if err := stream.Send(
				&operation_pb.FloorCeilingResponse{
					FloorValue:   math.Floor(request.GetNumber()),
					CeilingValue: math.Ceil(request.GetNumber()),
				}); err != nil {
				log.Fatalf("Failed to send floor_value & ceiling_value using stream. -%v", err)
			}
		}
	}
}
