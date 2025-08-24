// 代码生成时间: 2025-08-25 07:57:52
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/timestamppb"
)

// Constants for the system performance monitoring tool
const (
    // Address to listen on
    addr = "localhost:50051"
)

// SystemMetrics represents the metrics of the system
type SystemMetrics struct {
    Timestamp    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3"`
    CPUUsage     float32 `protobuf:"fixed32,2,opt,name=CPUUsage,proto3"`
    MemoryUsage  uint64 `protobuf:"varint,3,opt,name=MemoryUsage,proto3"`
    DiskUsage    uint64 `protobuf:"varint,4,opt,name=DiskUsage,proto3"`
}

// SystemPerformanceService defines the service for system performance monitoring
type SystemPerformanceService struct {
    // UnimplementedSystemPerformanceServiceServer can be embedded to have forward compatible implementations.
    unimplementedSystemPerformanceServiceServer
}

func (s *SystemPerformanceService) GetSystemMetrics(ctx context.Context, _ *emptypb.Empty) (*SystemMetrics, error) {
    // Here, we would implement the logic to gather system metrics
    // For demonstration purposes, we'll return mock data
    return &SystemMetrics{
        Timestamp: timestamppb.Now(),
        CPUUsage: 75.3,
        MemoryUsage: 2048, // in MB
        DiskUsage: 50000, // in MB
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Printf("Listening on %s", addr)

    // Create a new gRPC server
    s := grpc.NewServer()

    // Register the SystemPerformanceService with the gRPC server
    pb.RegisterSystemPerformanceServiceServer(s, &SystemPerformanceService{})

    // Set up a channel to listen for an interrupt or terminate signal from the OS.
    intCh := make(chan os.Signal, 1)
    signal.Notify(intCh, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-intCh
        s.GracefulStop()
    }()

    // Start the gRPC server
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// The following code is for the Protocol Buffers definition and should be placed in a .proto file
// For demonstration purposes, it is included here as comments

// syntax = "proto3\);\
//
// package pb;
//
// import "google/protobuf/timestamp.proto";
// import "google/protobuf/empty.proto";
//
// message SystemMetrics {
//     
//     google.protobuf.Timestamp timestamp = 1;
//     float32 CPUUsage = 2;
//     uint64 MemoryUsage = 3;
//     uint64 DiskUsage = 4;
// }
//
// service SystemPerformanceService {
//     rpc GetSystemMetrics(google.protobuf.Empty) returns (SystemMetrics);
// }

/*
This code represents a gRPC service that can be used for system performance monitoring.
It defines a service that, when called, returns a SystemMetrics message containing CPU usage,
memory usage, and disk usage. The service runs as a server, listening on localhost:50051,
and can be gracefully stopped with an interrupt signal.
*/