// 代码生成时间: 2025-10-09 02:39:29
// Package main implements a simple gRPC server with a defect tracking service.
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/encoding/protojson"
    "google.golang.org/protobuf/proto"
)

// Defect represents a defect in the system.
type Defect struct {
    Id          string
    Description string
    Status      string
}

// DefectServiceServer is the server API for DefectService service.
type DefectServiceServer struct{}

// AddDefect adds a new defect to the system.
func (s *DefectServiceServer) AddDefect(ctx context.Context, in *Defect) (*Defect, error) {
    if in.Id == "" || in.Description == "" || in.Status == "" {
        return nil, fmt.Errorf("defect id, description, and status cannot be empty")
    }
    // TODO: Implement the logic to add the defect to the system.
    fmt.Println("Defect added: ", in)
    return in, nil
}

// UpdateDefect updates an existing defect in the system.
func (s *DefectServiceServer) UpdateDefect(ctx context.Context, in *Defect) (*Defect, error) {
    if in.Id == "" || in.Description == "" || in.Status == "" {
        return nil, fmt.Errorf("defect id, description, and status cannot be empty")
    }
    // TODO: Implement the logic to update the defect in the system.
    fmt.Println("Defect updated: ", in)
    return in, nil
}

// GetDefect retrieves a defect by its ID.
func (s *DefectServiceServer) GetDefect(ctx context.Context, in *Defect) (*Defect, error) {
    if in.Id == "" {
        return nil, fmt.Errorf("defect id cannot be empty")
    }
    // TODO: Implement the logic to retrieve the defect by ID.
    fmt.Println("Defect retrieved: ", in)
    return in, nil
}

// DeleteDefect deletes a defect by its ID.
func (s *DefectServiceServer) DeleteDefect(ctx context.Context, in *Defect) (*Empty, error) {
    if in.Id == "" {
        return nil, fmt.Errorf("defect id cannot be empty")
    }
    // TODO: Implement the logic to delete the defect by ID.
    fmt.Println("Defect deleted: ", in)
    return &Empty{}, nil
}

// Empty represents an empty response.
type Empty struct{}

func main() {
   lis, err := net.Listen("tcp", ":50051")
   if err != nil {
       log.Fatalf("failed to listen: %v", err)
   }
   defer lis.Close()
   
   // Create a gRPC server.
   srv := grpc.NewServer()
   
   // Register the DefectServiceServer.
   // The DefectServiceServer must have a method corresponding to each method in the service.
   RegisterDefectServiceServer(srv, &DefectServiceServer{})
   
   // Register reflection service on gRPC server.
   reflection.Register(srv)
   
   fmt.Println("Server listening on port 50051")
   
   // Block and serve requests.
   if err := srv.Serve(lis); err != nil {
       log.Fatalf("failed to serve: %v", err)
   }
}
