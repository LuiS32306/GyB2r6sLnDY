// 代码生成时间: 2025-08-07 00:12:09
package main

import (
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/anypb"
)

// 定义消息通知的服务
type NotificationService struct {
    // 存储消息的map，模拟数据库存储
    messages map[string]string
}
inits := func() {
    return map[string]string{
        "message1": "Hello, this is a test message.",
        "message2": "This is another test message.",
    }
}

// 实现gRPC服务的方法
type notificationServiceServer struct {
    NotificationService
}
in // 发送消息的方法
func (s *notificationServiceServer) SendNotification(ctx grpc.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
    // 模拟消息存储
    s.messages[req.MessageId] = req.MessageContent

    // 模拟发送消息
    fmt.Printf("Sending message: %s\
", req.MessageContent)
    return &pb.SendNotificationResponse{Success: true}, nil
}
in // 获取消息的方法
func (s *notificationServiceServer) GetNotification(ctx grpc.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
    message, exists := s.messages[req.MessageId]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "Message not found")
    }
    return &pb.GetNotificationResponse{MessageContent: message}, nil
}
in // 初始化gRPC服务
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")

    // 创建gRPC服务器
    grpcServer := grpc.NewServer()

    // 注册服务和反射服务
    pb.RegisterNotificationServiceServer(grpcServer, &notificationServiceServer{NotificationService{messages: inits()})
    reflection.Register(grpcServer)

    // 启动服务器
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}