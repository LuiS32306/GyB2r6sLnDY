// 代码生成时间: 2025-08-21 10:45:44
package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// ErrorLogService 定义了一个错误日志收集服务
type ErrorLogService struct {
    // 文件存储器
    file *os.File
}

// NewErrorLogService 创建一个新的错误日志服务
func NewErrorLogService(filePath string) (*ErrorLogService, error) {
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return nil, err
    }
    return &ErrorLogService{file: file}, nil
}

// CollectErrorLog 实现收集错误日志的方法
func (els *ErrorLogService) CollectErrorLog(ctx context.Context, req *ErrorLogRequest) (*ErrorLogResponse, error) {
    // 记录错误日志到文件
    if _, err := els.file.WriteString(fmt.Sprintf("%s: %s
", time.Now().Format(time.RFC3339), req.GetErrorLog())); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to write error log: %v", err)
    }
    els.file.Sync() // 确保日志写入到磁盘
    return &ErrorLogResponse{Message: "Error log collected successfully"}, nil
}

// ErrorLogRequest 定义了错误日志请求的结构
type ErrorLogRequest struct {
    ErrorLog string `protobuf:"bytes,1,opt,name=error_log,json=errorLog"`
}

// ErrorLogResponse 定义了错误日志响应的结构
type ErrorLogResponse struct {
    Message string `protobuf:"bytes,1,opt,name=message"`
}

// server 是实现 ErrorLogService 的结构
type server struct {
    ErrorLogServiceServer
}

// RegisterServer 注册服务到gRPC服务器
func RegisterServer(s *grpc.Server, els *ErrorLogService) {
    ErrorLogServiceServer = els
    grpc.RegisterErrorLogServiceServer(s, &server{})
}

// main 函数是程序的入口点
func main() {
    const (
        address     = "localhost:50051" // 服务地址
        logFilePath = "error_logs.txt" // 错误日志文件路径
    )

    els, err := NewErrorLogService(logFilePath)
    if err != nil {
        log.Fatalf("failed to create error log service: %v", err)
    }
    defer els.file.Close()

    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Printf("Listening on %s...
", address)

    s := grpc.NewServer()
    RegisterServer(s, els)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
