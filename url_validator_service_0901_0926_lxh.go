// 代码生成时间: 2025-09-01 09:26:39
package main

import (
	"context"
	"fmt"
	"net/url"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// URLValidationService defines the service for validating URLs.
type URLValidationService struct{}

// ValidateURL will check if the provided URL is valid.
func (s *URLValidationService) ValidateURL(ctx context.Context, req *URLValidationRequest) (*URLValidationResponse, error) {
	// Check if the request is nil
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}

	// Parse the URL to check its validity
	parsedURL, err := url.Parse(req.GetUrl())
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid URL format")
	}

	// Check if the URL is well-formed
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		log.Printf("Invalid URL: %s", req.GetUrl())
		return nil, status.Errorf(codes.InvalidArgument, "URL must contain a scheme and a host")
	}

	// If all checks pass, return a valid response
	return &URLValidationResponse{IsValid: true}, nil
}

// URLValidationRequest holds the data needed for URL validation.
type URLValidationRequest struct {
	Url string `protobuf:"varint,1,opt,name=url,proto3" json:"url,omitempty"`
}

// URLValidationResponse holds the result of the URL validation.
type URLValidationResponse struct {
	IsValid bool `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
}

func main() {
	// gRPC server setup
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ls = grpc.NewServer()
	urlValidatorService := &URLValidationService{}
	pb.RegisterURLValidationServiceServer(ls, urlValidatorService)
	log.Printf("server listening at %v", lis.Addr())
	ls.Serve(nil)
}

// The following are the protobuf definitions for the service.
// They should be placed in a separate .proto file and compiled using the protoc tool.

// message URLValidationRequest {
//   string url = 1;
// }
//
// message URLValidationResponse {
//   bool is_valid = 1;
// }
//
// service URLValidationService {
//   rpc ValidateURL(URLValidationRequest) returns (URLValidationResponse);
// }