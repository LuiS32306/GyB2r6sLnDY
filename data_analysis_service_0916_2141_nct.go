// 代码生成时间: 2025-09-16 21:41:05
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
    "your_package_name/dataanalysis" // Replace with your actual package path
)

// Define the server
type DataAnalysisServer struct {}

// Define the method to sum numbers
func (s *DataAnalysisServer) Sum(ctx context.Context, req *dataanalysis.SumRequest) (*dataanalysis.SumResponse, error) {
    // Error handling
    if req == nil || req.GetNumbers() == nil {
        return nil, fmt.Errorf("invalid request")
    }

    // Calculate the sum
    sum := 0
    for _, num := range req.GetNumbers() {
        sum += num
    }

    // Return the sum
    return &dataanalysis.SumResponse{Sum: sum}, nil
}

// Define the method to calculate mean
func (s *DataAnalysisServer) Mean(ctx context.Context, req *dataanalysis.MeanRequest) (*dataanalysis.MeanResponse, error) {
    // Error handling
    if req == nil || req.GetNumbers() == nil {
        return nil, fmt.Errorf("invalid request")
    }

    // Calculate the mean
    sum := 0
    count := 0
    for _, num := range req.GetNumbers() {
        sum += num
        count++
    }

    // Check if count is zero to avoid division by zero
    if count == 0 {
        return nil, fmt.Errorf("cannot calculate mean for zero numbers")
    }

    // Return the mean
    return &dataanalysis.MeanResponse{Mean: float64(sum) / float64(count)}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    // Create a new gRPC server
    srv := grpc.NewServer()

    // Register the data analysis service on the server
    dataanalysis.RegisterDataAnalysisServer(srv, &DataAnalysisServer{})

    // Register reflection service on gRPC server.
    reflection.Register(srv)

    // Start the server
    if err := srv.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
