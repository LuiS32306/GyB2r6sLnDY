// 代码生成时间: 2025-09-24 01:09:31
package main

import (
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/structpb"
)

// ApiResponseFormatterService defines the gRPC service for API response formatting
type ApiResponseFormatterService struct{}

// FormatResponse formats the API response
func (s *ApiResponseFormatterService) FormatResponse(ctx grpc.Context, req *structpb.Struct) (*structpb.Struct, error) {
    // Check if the request is valid
    if req == nil {
        return nil, grpc.Errorf(codes.InvalidArgument, "Invalid request")
    }

    // Here you would add your own formatting logic
    // For demonstration purposes, we just return the received struct
    return req, nil
}

// server is used to implement api_response_formatter.ApiResponseFormatterServer
type server struct{
    api_response_formatter.UnimplementedApiResponseFormatterServer
}

// FormatResponse formats the API response
func (s *server) FormatResponse(ctx context.Context, req *api_response_formatter.FormatResponseRequest) (*api_response_formatter.FormatResponseResponse, error) {
    // Deserialize the request to a structpb.Struct
    var data structpb.Struct
    if err := jsonpb.Unmarshal(req.GetData(), &data); err != nil {
        return nil, grpc.Errorf(codes.InvalidArgument, "Failed to unmarshal request: %v", err)
    }

    // Call the formatting function
    formattedResponse, err := formatResponse(ctx, &data)
    if err != nil {
        return nil, err
    }

    // Serialize the formatted response back to jsonpb.Any
    serializedResponse, err := jsonpb.Marshal(formattedResponse)
    if err != nil {
        return nil, grpc.Errorf(codes.Internal, "Failed to serialize response: %v", err)
    }

    // Return the formatted response
    return &api_response_formatter.FormatResponseResponse{Data: serializedResponse}, nil
}

// formatResponse is a local function that formats the API response
func formatResponse(ctx context.Context, req *structpb.Struct) (*structpb.Struct, error) {
    // TODO: Implement your response formatting logic here
    // For now, we just return the request as is
    return req, nil
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("API Response Formatter service is listening on port 50051")
    
    s := grpc.NewServer()
    api_response_formatter.RegisterApiResponseFormatterServer(s, &server{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}