// 代码生成时间: 2025-07-31 02:44:18
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/emptypb"
)

// MathService provides a set of mathematical operations.
type MathService struct {
    // Embed an empty struct to allow embedding additional services.
    emptypb.UnimplementedEmptyServer
}

// CalculateSum calculates the sum of two numbers.
func (s *MathService) CalculateSum(ctx context.Context, in *SumRequest) (*SumResponse, error) {
    if in == nil || in.GetFirstNumber() == 0 && in.GetSecondNumber() == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "First number and second number cannot both be zero")
    }
    return &SumResponse{Result: in.GetFirstNumber() + in.GetSecondNumber()}, nil
}

// CalculateProduct calculates the product of two numbers.
func (s *MathService) CalculateProduct(ctx context.Context, in *ProductRequest) (*ProductResponse, error) {
    if in == nil || in.GetFirstNumber() == 0 && in.GetSecondNumber() == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "First number and second number cannot both be zero")
    }
    return &ProductResponse{Result: in.GetFirstNumber() * in.GetSecondNumber()}, nil
}

// Define the request and response messages for the sum operation.
type SumRequest struct {
    FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber" json:"first_number,omitempty"`
    SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber" json:"second_number,omitempty"`
}

type SumResponse struct {
    Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

// Define the request and response messages for the product operation.
type ProductRequest struct {
    FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber" json:"first_number,omitempty"`
    SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber" json:"second_number,omitempty"`
}

type ProductResponse struct {
    Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

// main is the entry point of the application.
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")

    s := grpc.NewServer()
    // Register the MathService server with the gRPC server.
    RegisterMathServiceServer(s, &MathService{})

    // Register reflection service on gRPC server.
    reflection.Register(s)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterMathServiceServer registers the MathServiceServer to the gRPC server.
func RegisterMathServiceServer(s *grpc.Server, srv *MathService) {
    // Register the MathService with the gRPC server.
    // This will make the server listen for incoming gRPC requests.
    RegisterMathServiceServer(s, srv)
}
