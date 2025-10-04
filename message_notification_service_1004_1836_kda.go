// 代码生成时间: 2025-10-04 18:36:45
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"path/to/your/message/notification/pb" // Replace with the actual path to your proto package
)

const (
	address     = ":50051" // The address to listen on
)

// server is used to implement pb.MessageNotificationServer.
type server struct {
	pb.UnimplementedMessageNotificationServer
}

// SayHello implements pb.MessageNotificationServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("Received: %v
", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageNotificationServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
