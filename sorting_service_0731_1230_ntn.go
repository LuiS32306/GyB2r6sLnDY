// 代码生成时间: 2025-07-31 12:30:53
package main

import (
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/proto"
)

// SortingService defines the sorting algorithm service.
type SortingService struct {
}

// Sort performs a bubble sort on the given slice of integers.
func (s *SortingService) Sort(ctx context.Context, in *pb.SortRequest) (*pb.SortResponse, error) {
    if in == nil || in.List == nil {
        return nil, status.Errorf(codes.InvalidArgument, "list cannot be empty")
    }

    list := make([]int, len(in.List))
    for i, v := range in.List {
        list[i] = int(v)
    }

    // Perform the bubble sort algorithm.
    for i := 0; i < len(list); i++ {
        for j := 0; j < len(list)-i-1; j++ {
            if list[j] > list[j+1] {
                // Swap the elements.
                list[j], list[j+1] = list[j+1], list[j]
            }
        }
    }

    // Convert the sorted list back to the protobuf message.
    sortedList := make([]*proto.Int32Value, len(list))
    for i, v := range list {
        sortedList[i] = &proto.Int32Value{Value: int32(v)}
    }

    return &pb.SortResponse{List: sortedList}, nil
}

// server is used to implement sortingservice.SortingServiceServer.
type server struct{
    sortingservice.UnimplementedSortingServiceServer
}

// NewServer creates a new SortingServiceServer.
func NewServer() *server {
    return &server{}
}

// Start starts the gRPC server and listens on the specified address.
func Start(address string) {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Printf("server listening at %s
", address)

    var opts []grpc.ServerOption
    s := grpc.NewServer(opts...)

    sortingservice.RegisterSortingServiceServer(s, NewServer())
    reflection.Register(s) // Register reflection service on gRPC server.
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    // Start the server on port 50051.
    Start(":50051")
}