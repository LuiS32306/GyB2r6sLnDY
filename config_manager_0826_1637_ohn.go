// 代码生成时间: 2025-08-26 16:37:47
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "os"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// 定义配置文件管理器服务
type ConfigManagerService struct {
    // 可以添加更多字段以存储配置信息
}

// 实现配置管理器服务接口
type ConfigManagerServer struct {
    ConfigManagerService
}

// GetConfig 方法用于获取配置文件的内容
func (s *ConfigManagerServer) GetConfig(ctx context.Context, req *ConfigRequest) (*ConfigResponse, error) {
    // 检查配置文件路径是否合法
    if _, err := os.Stat(req.FilePath); os.IsNotExist(err) {
        // 返回错误信息
        return nil, status.Errorf(codes.NotFound, "config file not found: %s", req.FilePath)
    }

    // 读取配置文件内容
    content, err := ioutil.ReadFile(req.FilePath)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to read config file: %s", err)
    }

    // 返回配置文件内容
    return &ConfigResponse{Content: string(content)}, nil
}

// SetConfig 方法用于设置配置文件的内容
func (s *ConfigManagerServer) SetConfig(ctx context.Context, req *SetConfigRequest) (*SetConfigResponse, error) {
    // 写入配置文件内容
    err := ioutil.WriteFile(req.FilePath, []byte(req.Content), 0644)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to write config file: %s", err)
    }

    // 返回成功响应
    return &SetConfigResponse{Success: true}, nil
}

// ConfigRequest 是获取配置请求的参数结构体
type ConfigRequest struct {
    FilePath string
}

// ConfigResponse 是获取配置响应的结果结构体
type ConfigResponse struct {
    Content string
}

// SetConfigRequest 是设置配置请求的参数结构体
type SetConfigRequest struct {
    FilePath string
    Content  string
}

// SetConfigResponse 是设置配置响应的结果结构体
type SetConfigResponse struct {
    Success bool
}

// main 函数启动 GRPC 服务器
func main() {
    // 监听端口
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建 GRPC 服务器
    s := grpc.NewServer()

    // 注册配置管理器服务
    configManagerServer := &ConfigManagerServer{}
    pb.RegisterConfigManagerServer(s, configManagerServer)

    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
