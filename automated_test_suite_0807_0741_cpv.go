// 代码生成时间: 2025-08-07 07:41:52
package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/test/bufconn"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Define your test service here
type TestServiceServer struct {}

// Define the service methods
func (s *TestServiceServer) TestMethod(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}
	return &TestResponse{Result: "Yes, it works!