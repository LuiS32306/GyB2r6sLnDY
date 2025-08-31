// 代码生成时间: 2025-08-31 20:59:34
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "google.golang.org/grpc"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/anypb"
)

type TextFileAnalyzerService struct {}

// Define the request message for the AnalyzeFile RPC
type AnalyzeFileRequest struct {
    FileName string
}

// Define the response message for the AnalyzeFile RPC
type AnalyzeFileResponse struct {
    Content string
    // Additional fields can be added here for more complex analysis results
}

// AnalyzeFile is the RPC method to analyze the content of a text file
func (s *TextFileAnalyzerService) AnalyzeFile(ctx context.Context, req *AnalyzeFileRequest) (*AnalyzeFileResponse, error) {
    // Check if the file exists
    if _, err := os.Stat(req.FileName); os.IsNotExist(err) {
        return nil, status.Error(codes.NotFound, "File not found")
    }

    // Read the content of the file
    content, err := ioutil.ReadFile(req.FileName)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "Failed to read file: %v", err)
    }

    // Convert the content to a string and return it in the response
    return &AnalyzeFileResponse{Content: string(content)}, nil
}

// RunServer starts the gRPC server and waits for incoming connections
func RunServer(port string, service *TextFileAnalyzerService) error {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        return status.Errorf(codes.Internal, "Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    RegisterTextFileAnalyzerServiceServer(s, service)
    if err := s.Serve(lis); err != nil {
        return status.Errorf(codes.Internal, "Failed to serve: %v", err)
    }
    return nil
}

// main function to start the server
func main() {
    port := ":50051" // Default port
    service := &TextFileAnalyzerService{}

    if err := RunServer(port, service); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}