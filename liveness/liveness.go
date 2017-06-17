package liveness

import "github.com/Zhanat87/go/grpc"

func CheckGrpcServerLiveness(in *grpc.EmptyRequest, stream grpc.GrpcService_CheckGrpcServerLivenessServer) error {
	livenessResponse := &grpc.LivenessResponse{
		Msg: "gRPC server works!",
	}

	if err := stream.Send(livenessResponse); err != nil {
		return err
	}
	return nil
}
