// 代码生成时间: 2025-09-21 09:57:52
package main

import (
    "context"
    "fmt"
    "net/url"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// URLValidationService provides methods to validate a URL
type URLValidationService struct {}

// ValidateURL checks if the provided URL is valid
func (s *URLValidationService) ValidateURL(ctx context.Context, req *URLValidationRequest) (*emptypb.Empty, error) {
    // Parse the URL to check its validity
    _, err := url.ParseRequestURI(req.Url)
    if err != nil {
        // Return a gRPC error with details
        return nil, status.Errorf(codes.InvalidArgument, "invalid URL: %v", err)
    }
    return &emptypb.Empty{}, nil
}

// URLValidationRequest represents the request for validating a URL
type URLValidationRequest struct {
    Url string
}

func main() {
    // Set up a gRPC server
    server := grpc.NewServer()
    fmt.Println("Starting gRPC server on port 50051")
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        fmt.Printf("Failed to listen: %v", err)
        return
    }
    defer lis.Close()

    // Register the URLValidationService with the gRPC server
    RegisterURLValidationServiceServer(server, &URLValidationService{})

    // Start serving gRPC requests
    if err := server.Serve(lis); err != nil {
        fmt.Printf("Failed to serve: %v", err)
    }
}

// RegisterURLValidationServiceServer registers the gRPC service with the server
func RegisterURLValidationServiceServer(s *grpc.Server, srv *URLValidationService) {
    // Register the service with the gRPC server
    RegisterURLValidationServiceServer(s, srv)
}
