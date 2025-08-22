// 代码生成时间: 2025-08-22 08:36:14
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "os"
    "strings"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// DatabaseMigrationService 定义了数据库迁移服务的接口
type DatabaseMigrationService interface {
    Migrate(context.Context, *MigrateRequest) (*MigrateResponse, error)
}

// MigrateRequest 定义了数据库迁移请求的结构
type MigrateRequest struct {
    // MigrationScriptPath 是迁移脚本的路径
    MigrationScriptPath string
}

// MigrateResponse 定义了数据库迁移响应的结构
type MigrateResponse struct {
    // Success 表示迁移是否成功
    Success bool
    // Message 提供迁移过程中的消息
    Message string
}

// databaseMigrationServer 实现了 DatabaseMigrationService 接口
type databaseMigrationServer struct{}

// Migrate 实现了数据库迁移服务
func (s *databaseMigrationServer) Migrate(ctx context.Context, req *MigrateRequest) (*MigrateResponse, error) {
    // 检查迁移脚本路径是否为空
    if req.MigrationScriptPath == "" {
        return nil, status.Errorf(codes.InvalidArgument, "migration script path cannot be empty")
    }

    // 读取迁移脚本文件
    scriptContent, err := os.ReadFile(req.MigrationScriptPath)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to read migration script: %v", err)
    }

    // 连接数据库并执行迁移脚本
    db, err := sql.Open("postgres", "user=youruser dbname=yourdb sslmode=disable")
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to connect to database: %v", err)
    }
    defer db.Close()

    // 执行迁移脚本
    if _, err := db.Exec(string(scriptContent)); err != nil {
        return nil, status.Errorf(codes.Internal, "failed to execute migration script: %v", err)
    }

    // 返回成功响应
    return &MigrateResponse{Success: true, Message: "Migration successful"}, nil
}

func main() {
    // 监听端口
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 创建gRPC服务器
    s := grpc.NewServer()

    // 注册数据库迁移服务
    DatabaseMigrationServiceServer(s, &databaseMigrationServer{})

    // 启动服务器
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
