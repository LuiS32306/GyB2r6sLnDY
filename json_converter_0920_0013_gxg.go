// 代码生成时间: 2025-09-20 00:13:16
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
    "encoding/json"
)

// Define the service
type JsonConverterServer struct{}

// Define the service methods
func (s *JsonConverterServer) Convert(ctx context.Context, req *JsonConverterRequest) (*JsonConverterResponse, error) {
    // Decode the JSON data from the request
    var jsonData interface{}
    err := json.Unmarshal(req.GetData(), &jsonData)
    if err != nil {
        return nil, err
    }

    // Marshal the JSON data back to bytes
    jsonDataBytes, err := json.Marshal(jsonData)
    if err != nil {
        return nil, err
    }

    // Return the response with the converted JSON data
    return &JsonConverterResponse{Data: jsonDataBytes}, nil
}

// Define the request and response messages
type JsonConverterRequest struct {
    Data []byte `protobuf:"varint,1,opt,name=data,proto3"`
}

type JsonConverterResponse struct {
    Data []byte `protobuf:"varint,1,opt,name=data,proto3"`
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    RegisterJsonConverterServer(s, &JsonConverterServer{})

    fmt.Println("Server listening on port 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterJsonConverterServer registers the JsonConverterServer to the gRPC server
func RegisterJsonConverterServer(s *grpc.Server, srv *JsonConverterServer) {
    RegisterJsonConverterServer(s, srv)
}
