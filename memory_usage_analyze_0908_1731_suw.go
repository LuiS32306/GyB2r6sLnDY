// 代码生成时间: 2025-09-08 17:31:51
package main

import (
    "fmt"
    "log"
    "net"
    "runtime"
    "google.golang.org/grpc"
)

// MemoryUsageService 定义了内存使用情况分析服务
type MemoryUsageService struct{}

// MemoryUsageResponse 定义了内存使用情况的响应结构
type MemoryUsageResponse struct {
    // HeapInuse 是当前堆内存使用量
    HeapInuse uint64
    // StackInuse 是当前栈内存使用量
    StackInuse uint64
    // MSpanInuse 是MSpan结构内存使用量
    MSpanInuse uint64
    // MCacheInuse 是MCache结构内存使用量
    MCacheInuse uint64
    // BuckHashSys 是BuckHashSys内存使用量
    BuckHashSys uint64
    // GCSys 是GC内存使用量
    GCSys uint64
    // OtherSys 是其他内存使用量
    OtherSys uint64
    // NextGC 是下一次GC的内存使用量
    NextGC uint64
    // LastGC 是上次GC的时间
    LastGC uint64 // nanoseconds
    // PauseTotalNs 是GC暂停总时间
    PauseTotalNs uint64
    // NumGC 是GC次数
    NumGC uint32
}

// GetMemoryUsage 实现了内存使用情况分析服务，返回当前内存使用情况
func (s *MemoryUsageService) GetMemoryUsage(_ *empty.Empty, stream MemoryUsageService_GetMemoryUsageServer) error {
    for {
        // 获取当前内存使用情况
        memStats := runtime.MemStats{}
        runtime.ReadMemStats(&memStats)

        // 构建响应
        response := MemoryUsageResponse{
            HeapInuse:    memStats.HeapInuse,
            StackInuse:   memStats.StackInuse,
            MSpanInuse:   memStats.MSpanInuse,
            MCacheInuse:  memStats.MCacheInuse,
            BuckHashSys:  memStats.BuckHashSys,
            GCSys:        memStats.GCSys,
            OtherSys:     memStats.OtherSys,
            NextGC:       memStats.NextGC,
            LastGC:       memStats.LastGC,
            PauseTotalNs: memStats.PauseTotalNs,
            NumGC:        memStats.NumGC,
        }

        // 发送响应
        if err := stream.Send(&response); err != nil {
            return err
        }

        // 每秒更新一次
        runtime.Gosched()
    }
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on :50051")

    // 创建gRPC服务器
    server := grpc.NewServer()

    // 注册服务
    RegisterMemoryUsageServiceServer(server, &MemoryUsageService{})

    // 启动服务
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
