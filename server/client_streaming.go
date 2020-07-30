package server

import (
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"io"
)

func (*Service) ComputeAverage(stream operation_pb.OperationService_ComputeAverageServer) error {
	var sum int64
	var count = 0
	for {
		if request, err := stream.Recv(); err == io.EOF {
			return stream.SendAndClose(
				&operation_pb.ComputeAverageResponse{
					Average: float64(sum) / float64(count),
				})
		} else if err != nil {
			fatalLogger("Failed to read the number from the stream. - %v", err)
		} else {
			sum += request.GetNumber()
			count++
		}
	}
}
