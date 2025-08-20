// 代码生成时间: 2025-08-20 10:45:14
// batch_rename_tool.go
// 批量文件重命名工具，使用GRPC框架实现服务端和客户端之间的通信。

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"google.golang.org/grpc"
)

// 文件重命名请求
type RenameRequest struct {
	OldName string // 原始文件名
	NewName string // 新文件名
}

// 文件重命名响应
type RenameResponse struct {
	Success bool   // 是否成功
	Message string // 信息
}

// 文件服务
type FileService struct{}

// Rename 方法用于批量重命名文件
func (s *FileService) Rename(ctx context.Context, req *RenameRequest) (*RenameResponse, error) {
	// 检查原始文件是否存在
	if _, err := os.Stat(req.OldName); os.IsNotExist(err) {
		return &RenameResponse{Success: false, Message: "Original file does not exist"}, nil
	}
	
	// 尝试重命名文件
	err := os.Rename(req.OldName, req.NewName)
	if err != nil {
		return &RenameResponse{Success: false, Message: "Failed to rename file"}, nil
	}
	
	return &RenameResponse{Success: true, Message: "File renamed successfully"}, nil
}

func main() {
	// 初始化GRPC服务端
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	
	fmt.Println("File Rename Service is running on port 50051")
	
	// 注册服务
	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, &FileService{})
	
	// 启动服务端
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
