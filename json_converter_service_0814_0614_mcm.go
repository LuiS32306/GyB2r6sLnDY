// 代码生成时间: 2025-08-14 06:14:29
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/proto"
    "google.golang.org/protobuf/types/known/anypb"
)

// JSONDataConverterService defines the service that converts JSON data to and from protobuf.
type JSONDataConverterService struct{}

// ConvertJSONToProtobuf converts JSON data to protobuf message.
func (s *JSONDataConverterService) ConvertJSONToProtobuf(ctx context.Context, req *JSONDataRequest) (*ProtobufResponse, error) {
    if req == nil || req.JsonData == "" {
        return nil, fmt.Errorf("invalid request")
    }

    var msg proto.Message
    switch req.MessageName {
    case "SomeMessage":
        msg = &SomeMessage{}
        // Add more message types as needed.
    default:
        return nil, fmt.Errorf("unsupported message type: %s", req.MessageName)
    }

    err := protojson.Unmarshal([]byte(req.JsonData), msg)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
    }

    any, err := anypb.NewProtoMsg(msg)
    if err != nil {
        return nil, fmt.Errorf("failed to create Any message: %w", err)
    }

    return &ProtobufResponse{Data: any}, nil
}

// ConvertProtobufToJSON converts protobuf message to JSON data.
func (s *JSONDataConverterService) ConvertProtobufToJSON(ctx context.Context, req *ProtobufRequest) (*JSONDataResponse, error)
{
    if req == nil || req.Data == nil {
        return nil, fmt.Errorf("invalid request")
    }

    var msg proto.Message
    switch req.MessageName {
    case "SomeMessage":
        msg = &SomeMessage{}
        // Add more message types as needed.
    default:
        return nil, fmt.Errorf("unsupported message type: %s", req.MessageName)
    }

    err := proto.Unmarshal(req.Data.Value, msg)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal protobuf: %w", err)
    }

    jsonData, err := protojson.Marshal(msg)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal to JSON: %w", err)
    }

    return &JSONDataResponse{JsonData: string(jsonData)}, nil
}

// Main function to start the gRPC server and listen for incoming requests.
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    // Register your service with the server.
    RegisterJSONDataConverterServiceServer(s, &JSONDataConverterService{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Define the request and response messages for the service.
type JSONDataRequest struct {
    JsonData  string `protobuf:"bytes,1,opt,name=json_data,json=jsonData"`
    MessageName string `protobuf:"bytes,2,opt,name=message_name,json=messageName"`
}

type ProtobufResponse struct {
    Data *anypb.Any `protobuf:"bytes,1,opt,name=data"`
}

type ProtobufRequest struct {
    Data      *anypb.Any `protobuf:"bytes,1,opt,name=data"`
    MessageName string `protobuf:"bytes,2,opt,name=message_name,json=messageName"`
}

type JSONDataResponse struct {
    JsonData string `protobuf:"bytes,1,opt,name=json_data,json=jsonData"`
}

// Define the service in a protobuf file and then generate the Go code using protoc.
/*
syntax = "proto3";

package jsonconverter;

service JSONDataConverterService {
    rpc ConvertJSONToProtobuf(JSONDataRequest) returns (ProtobufResponse);
    rpc ConvertProtobufToJSON(ProtobufRequest) returns (JSONDataResponse);
}

message JSONDataRequest {
    string json_data = 1;
    string message_name = 2;
}

message ProtobufResponse {
    google.protobuf.Any data = 1;
}

message ProtobufRequest {
    google.protobuf.Any data = 1;
    string message_name = 2;
}

message JSONDataResponse {
    string json_data = 1;
}
*/
