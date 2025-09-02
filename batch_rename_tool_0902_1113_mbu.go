// 代码生成时间: 2025-09-02 11:13:15
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/emptypb"
)

// 文件重命名请求
type RenameRequest struct {
    Source string // 源文件路径
    Target string // 目标文件路径
}

// 文件重命名响应
type RenameResponse struct {
    Success bool   // 重命名是否成功
    Message string // 重命名结果消息
}

// 文件重命名服务
type RenameService struct {
    UnimplementedRenameServer
}

// Rename 实现批量文件重命名
func (s *RenameService) Rename(ctx context.Context, req *RenameRequest) (*RenameResponse, error) {
    source := req.GetSource()
    target := req.GetTarget()
    
    // 检查源文件是否存在
    if _, err := os.Stat(source); os.IsNotExist(err) {
        return &RenameResponse{
            Success: false,
            Message: "Source file does not exist.",
        }, nil
    }
    
    // 检查目标文件是否已存在
    if _, err := os.Stat(target); !os.IsNotExist(err) {
        return &RenameResponse{
            Success: false,
            Message: "Target file already exists.",
        }, nil
    }
    
    // 重命名文件
    if err := os.Rename(source, target); err != nil {
        return &RenameResponse{
            Success: false,
            Message: fmt.Sprintf("Failed to rename file: %v", err),
        }, nil
    }
    
    return &RenameResponse{
        Success: true,
        Message: "File renamed successfully.",
    }, nil
}

// 启动gRPC服务器
func startGRPCServer() {
    lis, err := grpc.NetListener("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    defer lis.Close()
    
    s := grpc.NewServer()
    rename.RegisterRenameServer(s, &RenameService{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    startGRPCServer()
}
