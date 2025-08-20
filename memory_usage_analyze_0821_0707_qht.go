// 代码生成时间: 2025-08-21 07:07:09
package main

import (
    "fmt"
    "net"
    "os"
    "runtime"
    "runtime/pprof"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

// MemoryUsageService provides a service for analyzing memory usage.
type MemoryUsageService struct{}

// GetMemoryUsage returns the current memory usage statistics.
func (m *MemoryUsageService) GetMemoryUsage(ctx context.Context, req *MemoryUsageRequest) (*MemoryUsageResponse, error) {
    // Start profiling the memory usage.
    f, err := os.Create("profile.mem")
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to create profile file: %v", err)
    }
    defer f.Close()

    if err := pprof.WriteHeapProfile(f); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to write heap profile: %v", err)
    }

    // Calculate memory usage statistics.
    stats := &MemoryUsageResponse{
        HeapInUse: runtime.MemStats.HeapInuse,
        HeapReleased: runtime.MemStats.HeapReleased,
    }

    return stats, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")

    s := grpc.NewServer()
    memoryUsageService := &MemoryUsageService{}
    memoryusage.RegisterMemoryUsageServer(s, memoryUsageService)
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// MemoryUsageRequest is a request for memory usage data.
type MemoryUsageRequest struct{}

// MemoryUsageResponse is a response containing memory usage statistics.
type MemoryUsageResponse struct {
    HeapInUse       uint64 `protobuf:"varint,1,opt,name=heap_in_use,json=heapInUse" json:"heap_in_use,omitempty"`
    HeapReleased    uint64 `protobuf:"varint,2,opt,name=heap_released,json=heapReleased" json:"heap_released,omitempty"`
}

// MemoryUsageServer is the server API for MemoryUsage service.
type MemoryUsageServer interface {
    GetMemoryUsage(context.Context, *MemoryUsageRequest) (*MemoryUsageResponse, error)
}

// UnimplementedMemoryUsageServer can be embedded to have forward compatible implementations.
type UnimplementedMemoryUsageServer struct{}

func (*UnimplementedMemoryUsageServer) GetMemoryUsage(ctx context.Context, req *MemoryUsageRequest) (*MemoryUsageResponse, error) {
    return nil, status.Errorf(codes.Unimplemented, "method GetMemoryUsage not implemented")
}
