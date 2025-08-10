// 代码生成时间: 2025-08-10 12:38:05
package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CacheService is the service that will be exposed via gRPC.
type CacheService struct {
	sync.RWMutex
	cache map[string]string
}
pm
// NewCacheService creates a new instance of CacheService with an empty cache.
func NewCacheService() *CacheService {
	return &CacheService{
		cache: make(map[string]string),
	}
}

// Set adds or updates the value for a key in the cache.
func (s *CacheService) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	s.Lock()
	defer s.Unlock()
	s.cache[req.Key] = req.Value
	return &SetResponse{}, nil
}

// Get retrieves the value for a key from the cache.
func (s *CacheService) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	s.RLock()
	defer s.RUnlock()
	value, exists := s.cache[req.Key]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "key not found")
	}
	return &GetResponse{Value: value}, nil
}

// Clear removes the key from the cache.
func (s *CacheService) Clear(ctx context.Context, req *ClearRequest) (*ClearResponse, error) {
	s.Lock()
	defer s.Unlock()
	delete(s.cache, req.Key)
	return &ClearResponse{}, nil
}

// StartServer starts the gRPC server with the given address and service implementation.
func StartServer(addr string, service *CacheService) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("Listening on %s
", addr)
	gls := grpc.NewServer()
	pb.RegisterCacheServiceServer(gls, service)
	gls.Serve(lis)
}

func main() {
	cacheService := NewCacheService()
	StartServer(":50051", cacheService)
}

// Below are the gRPC message definitions.

type SetRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key"`
	Value string `protobuf:"bytes,2,opt,name=value"`
}

type SetResponse struct{}

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key"`
}

type GetResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value"`
}

type ClearRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key"`
}

type ClearResponse struct{}

// Below are the gRPC service definitions.

service CacheService {
  rpc Set(SetRequest) returns (SetResponse) {}
  rpc Get(GetRequest) returns (GetResponse) {}
  rpc Clear(ClearRequest) returns (ClearResponse) {}
}
