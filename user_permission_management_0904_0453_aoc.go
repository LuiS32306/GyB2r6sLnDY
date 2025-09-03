// 代码生成时间: 2025-09-04 04:53:25
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection" 
    "google.golang.org/protobuf/types/known/emptypb" 
)

// Define the permissions that can be assigned to users.
const (
    PERMISSION_READ  = "read"
    PERMISSION_WRITE = "write"
)

// UserPermission defines the permissions for a user.
type UserPermission struct {
    UserID  string
    Perms   []string
}

// UserPermissionService provides methods for managing user permissions.
type UserPermissionService struct {
    // Maps user IDs to their permissions.
    permissions map[string][]string
}

// NewUserPermissionService creates a new service with an empty permissions map.
func NewUserPermissionService() *UserPermissionService {
    return &UserPermissionService{
        permissions: make(map[string][]string),
    }
}

// AddPermission assigns a permission to a user.
func (s *UserPermissionService) AddPermission(ctx context.Context, req *UserPermission) (*emptypb.Empty, error) {
    if _, exists := s.permissions[req.UserID]; !exists {
        s.permissions[req.UserID] = []string{}
    }
    for _, perm := range req.Perms {
        s.permissions[req.UserID] = append(s.permissions[req.UserID], perm)
    }
    return &emptypb.Empty{}, nil
}

// RemovePermission removes a permission from a user.
func (s *UserPermissionService) RemovePermission(ctx context.Context, req *UserPermission) (*emptypb.Empty, error) {
    if perms, exists := s.permissions[req.UserID]; exists {
        for i, perm := range perms {
            if perm == req.Perms[0] {
                s.permissions[req.UserID] = append(perms[:i], perms[i+1:]...)
                break
            }
        }
    }
    return &emptypb.Empty{}, nil
}

// CheckPermission checks if a user has a specific permission.
func (s *UserPermissionService) CheckPermission(ctx context.Context, req *UserPermission) (*emptypb.Empty, error) {
    if perms, exists := s.permissions[req.UserID]; exists {
        for _, perm := range perms {
            if perm == req.Perms[0] {
                return &emptypb.Empty{}, nil
            }
        }
    }
    return nil, fmt.Errorf("permission denied")
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    service := NewUserPermissionService()
    // Register the service with the GRPC server.
    // RegisterUserPermissionServiceServer(s, service)
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}