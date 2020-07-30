package client

import (
	"context"
	"github.com/aka-achu/go-gRPC/models/operation_pb"
	"io"
	"time"
)

func FloorCeiling(c operation_pb.OperationServiceClient) {

	if stream, err := c.FloorCeiling(context.Background()); err != nil {
		fatalLogger("Failed to create a stream connection. -%v", err)
	} else {
		var wait = make(chan struct{})
		go func() {
			for _, number := range []float64{1.23, 2.56, 9.36, 7.42, 5.11, 0.88} {
				if err := stream.Send(
					&operation_pb.FloorCeilingRequest{
						Number: number,
					}); err != nil {
					fatalLogger("Failed to send number in the stream channel. -%v", err)
				}
				time.Sleep(time.Second)
			}
			_ = stream.CloseSend()
		}()

		go func() {
			for {
				if response, err := stream.Recv(); err == io.EOF {
					break
				} else if err != nil {
					fatalLogger("Failed to receive the floor_value and ceiling_value from the stream. -%v", err)
				} else {
					printer("Response from server- Floor Value:%f Ceiling Value:%f",
						response.GetFloorValue(), response.GetCeilingValue())
				}
			}
			close(wait)
		}()

		<-wait
	}

}
