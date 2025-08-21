// 代码生成时间: 2025-08-22 00:15:09
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/grpc/status"
)

// NetworkConnectionCheckerService defines the gRPC service for checking network connections.
type NetworkConnectionCheckerService struct{}

// CheckConnection is a gRPC method that checks if a network connection is available.
func (n *NetworkConnectionCheckerService) CheckConnection(ctx context.Context, req *CheckConnectionRequest) (*CheckConnectionResponse, error) {
    // Try to establish a connection to the host specified in the request.
    conn, err := net.DialTimeout("tcp", req.GetHost(), 5*time.Second)
    if err != nil {
        // If an error occurs, return the error status.
        return nil, status.Errorf(codes.Unavailable, "could not connect to host: %v", err)
    }
    defer conn.Close()

    // If the connection is successful, return a success status.
    return &CheckConnectionResponse{
        Success: true,
    }, nil
}

// CheckConnectionRequest is the request message for the CheckConnection method.
type CheckConnectionRequest struct {
    Host string `protobuf:"bytes,1,opt,name=host"`
}

// CheckConnectionResponse is the response message for the CheckConnection method.
type CheckConnectionResponse struct {
    Success bool `protobuf:"varint,1,opt,name=success"`
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port :50051")

    s := grpc.NewServer()
    RegisterNetworkConnectionCheckerServiceServer(s, &NetworkConnectionCheckerService{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterNetworkConnectionCheckerServiceServer registers the gRPC service.
func RegisterNetworkConnectionCheckerServiceServer(s *grpc.Server, srv *NetworkConnectionCheckerService) {
    pb.RegisterNetworkConnectionCheckerServiceServer(s, srv)
}