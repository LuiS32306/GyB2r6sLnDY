// 代码生成时间: 2025-08-03 03:39:45
package main

import (
    "fmt"
    "google.golang.org/grpc"
    "net"
    "sort"
)

// SortingAlgorithmService 定义了排序算法的服务接口
type SortingAlgorithmService struct{}

// SortNumbers 是 SortingAlgorithmService 提供的 RPC 方法，用于排序传入的数字列表
func (s *SortingAlgorithmService) SortNumbers(ctx context.Context, in *pb.SortRequest) (*pb.SortResponse, error) {
    // 将接收到的数字列表转换为 slice
    numbers := make([]int, len(in.Numbers))
    for i, num := range in.Numbers {
        numbers[i] = int(num)
    }

    // 使用 Go 标准库对数字进行排序
    sort.Ints(numbers)

    // 创建响应消息，将排序后的数字列表转换回 protobuf 消息
    sortedNumbers := make([]*pb.Int, len(numbers))
    for i, num := range numbers {
        sortedNumbers[i] = &pb.Int{Value: int32(num)}
    }
    return &pb.SortResponse{Numbers: sortedNumbers}, nil
}

// server 是一个包含 SortingAlgorithmService 和 gRPC 服务的 gRPC 服务器
type server struct{
    pb.UnimplementedSortingAlgorithmServiceServer
}

// RegisterServer 注册排序算法服务到 gRPC 服务器
func RegisterServer(s *grpc.Server, service *SortingAlgorithmService) {
    pb.RegisterSortingAlgorithmServiceServer(s, service)
}

// main 函数中，我们启动一个 gRPC 服务器，并注册排序算法服务
func main() {
    fmt.Println("Starting Sorting Algorithm Service...")
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        fmt.Println("Failed to listen: ", err)
        return
    }
    fmt.Println("Listening on port 50051...
")
    s := grpc.NewServer()
    RegisterServer(s, &SortingAlgorithmService{})
    if err := s.Serve(lis); err != nil {
        fmt.Println("Failed to serve: ", err)
        return
    }
}