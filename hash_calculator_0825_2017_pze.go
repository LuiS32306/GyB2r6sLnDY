// 代码生成时间: 2025-08-25 20:17:16
package main

import (
    "fmt"
    "hash"
    "log"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    pb "path/to/your/protobuf/definitions" // Replace with the actual path to your protobuf definitions
)

// server is used to implement hashcalculator.HashCalculatorServer.
type server struct {
    pb.UnimplementedHashCalculatorServer
}

// CalculateHash implements hashcalculator.HashCalculatorServer
func (s *server) CalculateHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
    // Choose the hash function based on the request
    switch in.HashType {
    case pb.HashType_SHA256:
        result := hash.HashString(in.StringToHash, "sha256")
        return &pb.HashResponse{Hash: result}, nil
    case pb.HashType_MD5:
        result := hash.HashString(in.StringToHash, "md5")
        return &pb.HashResponse{Hash: result}, nil
    // Add more hash types as needed
    default:
        return nil, fmt.Errorf("unsupported hash type %v", in.HashType)
    }
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")
    
    s := grpc.NewServer()
    pb.RegisterHashCalculatorServer(s, &server{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}