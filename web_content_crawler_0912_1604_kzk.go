// 代码生成时间: 2025-09-12 16:04:51
This service provides a gRPC interface to fetch web page content.
It includes error handling, proper documentation, and follows GoLang best practices.

Author: Your Name
Date: 2023-04-01
*/

package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// Define the gRPC service
type WebContentCrawlerService struct {}

// Define the function to fetch web content
func (s *WebContentCrawlerService) FetchWebContent(ctx context.Context, in *WebContentRequest) (*WebContentResponse, error) {
    // Check for invalid URL
    if in.Url == "" {
        return nil, status.Error(codes.InvalidArgument, "Invalid URL")
    }

    // Fetch web content
    resp, err := http.Get(in.Url)
    if err != nil {
        return nil, status.Error(codes.Unavailable, "Failed to fetch web content")
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, status.Error(codes.Internal, "Failed to read response body")
    }

    // Create the response
    return &WebContentResponse{Content: string(body)}, nil
}

// Define the gRPC server
func RunServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    RegisterWebContentCrawlerServiceServer(grpcServer, &WebContentCrawlerService{})
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Define the gRPC request and response messages
type WebContentRequest struct {
    Url string
}

type WebContentResponse struct {
    Content string
}

func main() {
    RunServer()
}
