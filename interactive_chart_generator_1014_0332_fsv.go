// 代码生成时间: 2025-10-14 03:32:26
// interactive_chart_generator.go
// 这个程序是一个交互式图表生成器，使用GOLANG和GRPC框架实现。

package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "log"
    "time"
    "math/rand"

    "google.golang.org/protobuf/types/known/timestamppb"
    "google.golang.org/protobuf/types/known/emptypb"

    pb "interactive_chart_generator/pb" // 假设pb是包含gRPC服务定义的包
)

// ChartServiceServer 是gRPC服务的服务器实现
type ChartServiceServer struct{}

// GenerateChart 是gRPC服务的一个方法，用于生成交互式图表
func (s *ChartServiceServer) GenerateChart(ctx context.Context, in *pb.ChartRequest) (*pb.ChartResponse, error) {
    // 检查请求参数是否合法
    if in == nil || in.GetTitle() == "" {
        return nil, grpc.Errorf(codes.InvalidArgument, "missing chart title")
    }

    // 模拟图表数据生成
    var data []float64
    for i := 0; i < 10; i++ {
        data = append(data, rand.Float64())
    }

    // 创建响应
    return &pb.ChartResponse{
        Title:   in.GetTitle(),
        Data:    data,
        GeneratedAt: timestamppb.Now(),
    }, nil
}

// main 是程序入口点
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port :50051")

    server := grpc.NewServer()
    pb.RegisterChartServiceServer(server, &ChartServiceServer{})

    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
