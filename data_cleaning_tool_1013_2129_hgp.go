// 代码生成时间: 2025-10-13 21:29:50
package main

import (
    "context"
    "fmt"
    "log"
    "google.golang.org/grpc"
)

// DataCleaningService 定义gRPC服务
type DataCleaningService struct {
}

// CleanData 实现数据清洗接口
func (s *DataCleaningService) CleanData(ctx context.Context, in *DataCleaningRequest) (*DataCleaningResponse, error) {
    // 这里可以添加数据清洗和预处理的逻辑
    // 例如，去除空格，过滤无效数据等

    // 示例：去除空白字符
    cleanedData := strings.TrimSpace(in.GetData())

    // 检查数据是否有效，这里简单示例为非空检查
    if len(cleanedData) == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "received empty data")
    }

    // 返回清洗后的数据
    return &DataCleaningResponse{Data: cleanedData}, nil
}

// Main 方法实现gRPC服务器的启动和运行
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")

    s := grpc.NewServer()
    pb.RegisterDataCleaningServiceServer(s, &DataCleaningService{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// DataCleaningRequest 请求消息定义
type DataCleaningRequest struct {
    Data string `protobuf:"bytes,1,opt,name=data,proto3"`
}

// DataCleaningResponse 响应消息定义
type DataCleaningResponse struct {
    Data string `protobuf:"bytes,1,opt,name=data,proto3"`
}

// DataCleaningServiceServer 定义gRPC服务服务器接口
type DataCleaningServiceServer interface {
    CleanData(context.Context, *DataCleaningRequest) (*DataCleaningResponse, error)
}

// RegisterDataCleaningServiceServer 注册gRPC服务
func RegisterDataCleaningServiceServer(s *grpc.Server, srv DataCleaningServiceServer) {
    pb.RegisterDataCleaningServiceServer(s, srv)
}
