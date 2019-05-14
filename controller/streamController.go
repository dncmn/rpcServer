package controller

import (
	"fmt"
	"io"
	"log"
	pb "rpcServer/pb"
)

// 服务端流rpc
func (s *StreamService) List(r *pb.StreamRequest, stream pb.StreamService_ListServer) error {
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.StreamResponse{
			Pt: &pb.StreamPoint{
				Name:  r.Pt.Name,
				Value: r.Pt.Value + int32(n),
			},
		})
		if err != nil {
			return err
		}
		fmt.Println(n)
	}
	return nil
}

// 客户端流rpc
func (s *StreamService) Record(stream pb.StreamService_RecordServer) error {
	for {
		r, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.StreamResponse{Pt: &pb.StreamPoint{Name: "grpc Stream Server: Record", Value: 1}})
		}
		if err != nil {
			return err
		}
		log.Printf("stream.Recv pt.Name=%v,pt.Value=%v\n", r.Pt.Name, r.Pt.Value)
	}
	return nil
}

// 双向流rpc
func (s *StreamService) Route(stream pb.StreamService_RouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.StreamResponse{Pt: &pb.StreamPoint{
			Name:  "gRPC Stream Client:Route",
			Value: int32(n),
		}})
		if err != nil {
			return err
		}

		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}
		n++
		log.Printf("stream.Recv pt.Name=%v,pt.Value=%v\n", r.Pt.Name, r.Pt.Value)
	}
	return nil
}
