// 代码生成时间: 2025-10-10 17:18:16
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
    "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
    "github.com/grpc-ecosystem/go-grpc-middleware/tags/zap"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// HomeworkService 作业服务定义
type HomeworkService struct{}

// Homework 作业定义
type Homework struct {
    Id      int    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
    Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
    Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
    DueDate string `protobuf:"bytes,4,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
}

// HomeworkServiceServer 作业服务服务器
type HomeworkServiceServer struct{
    HomeworkService
}

// CreateHomework 创建作业
func (s *HomeworkServiceServer) CreateHomework(ctx context.Context, req *CreateHomeworkRequest) (*Homework, error) {
    // 这里添加创建作业的逻辑
    // 为了演示，我们直接返回请求中的作业
    return req.Homework, nil
}

// GetHomework 获取作业
func (s *HomeworkServiceServer) GetHomework(ctx context.Context, req *GetHomeworkRequest) (*Homework, error) {
    // 这里添加获取作业的逻辑
    // 为了演示，我们直接返回请求中的作业
    return req.Homework, nil
}

// UpdateHomework 更新作业
func (s *HomeworkServiceServer) UpdateHomework(ctx context.Context, req *UpdateHomeworkRequest) (*Homework, error) {
    // 这里添加更新作业的逻辑
    // 为了演示，我们直接返回请求中的作业
    return req.Homework, nil
}

// DeleteHomework 删除作业
func (s *HomeworkServiceServer) DeleteHomework(ctx context.Context, req *DeleteHomeworkRequest) (*emptypb.Empty, error) {
    // 这里添加删除作业的逻辑
    // 为了演示，我们直接返回空响应
    return &emptypb.Empty{}, nil
}

// RegisterHomeworkServiceServer 注册作业服务服务器
func RegisterHomeworkServiceServer(server *grpc.Server, service HomeworkService) {
    RegisterHomeworkServiceServer(server, &HomeworkServiceServer{HomeworkService: service})
}

// NewHomeworkService 创建作业服务
func NewHomeworkService() HomeworkService {
    return HomeworkService{}
}

// RunServer 运行服务器
func RunServer(address string, service HomeworkService) {
    lis, err := net.Listen("tcp", address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
        grpc_zap.UnaryServerInterceptor(zap.L()),
        grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoveryHandler)),
    )))
    RegisterHomeworkServiceServer(server, service)
    reflection.Register(server)
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// recoveryHandler 错误处理函数
func recoveryHandler(p panic.Recovered, i interface{}) error {
    return status.Errorf(codes.Internal, fmt.Sprintf("panic triggered: %v", p))
}

// main 程序入口
func main() {
    address := ":50051"
    service := NewHomeworkService()
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    go func() {
        c := make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        <-c
        log.Println("shutting down gRPC server...")
        cancel()
        if err := ctx.Err(); err != nil {
            log.Fatalf("failed to shutdown server: %v", err)
        }
    }()
    RunServer(address, service)
}
