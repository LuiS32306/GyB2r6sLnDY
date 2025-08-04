// 代码生成时间: 2025-08-04 18:16:54
package main

import (
    "fmt"
    "hash"
    "log"
    "net"
    "golang.org/x/crypto/sha3"
    "google.golang.org/grpc"
)

// HashValueRequest is the request message containing the data to hash.
type HashValueRequest struct {
    Data string
}

// HashValueResponse is the response message containing the hash value.
type HashValueResponse struct {
    HashValue string
}

// HashService provides a service for hashing data.
type HashService struct{}

// ComputeHash calculates the SHA-3 hash of the provided data.
func (s *HashService) ComputeHash(ctx context.Context, req *HashValueRequest) (*HashValueResponse, error) {
    if req.Data == "" {
        return nil, status.Errorf(codes.InvalidArgument, "empty data not allowed")
    }

    hashValue := ComputeSHA3Hash(req.Data)
    return &HashValueResponse{HashValue: hashValue}, nil
}

// ComputeSHA3Hash computes the SHA-3 hash of the provided data.
func ComputeSHA3Hash(data string) string {
    hash := sha3.New256()
    _, err := hash.Write([]byte(data))
    if err != nil {
        log.Fatalf("Failed to compute hash: %v", err)
    }
    return fmt.Sprintf("%x", hash.Sum(nil))
}

// server is used to implement HashServiceServer.
type server struct{
    HashServiceServer
}

// ComputeHash implements HashServiceServer
func (s *server) ComputeHash(ctx context.Context, req *HashValueRequest) (*HashValueResponse, error) {
    return hashService.ComputeHash(ctx, req)
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Hash service listening on port :50051")

    grpcServer := grpc.NewServer()
    hashService := &HashService{}
    pb.RegisterHashServiceServer(grpcServer, &server{HashServiceServer: hashService})

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
