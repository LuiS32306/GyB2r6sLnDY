// 代码生成时间: 2025-08-02 19:35:19
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/encoding/protojson"
)

// ThemeService defines the structure for the gRPC service.
type ThemeService struct {
    // Additional fields can be added for service logic
}

// ThemeRequest message for theme switch request
type ThemeRequest struct {
    Theme string
}

// ThemeResponse message for theme switch response
type ThemeResponse struct {
    Success bool
}

// ThemeServiceServer defines the server interface.
type ThemeServiceServer interface {
    ChangeTheme(context.Context, *ThemeRequest) (*ThemeResponse, error)
}

// ThemeServer implements ThemeServiceServer.
type ThemeServer struct {
    ThemeServiceServer
}

// ChangeTheme changes the theme based on the request.
func (t *ThemeServer) ChangeTheme(ctx context.Context, req *ThemeRequest) (*ThemeResponse, error) {
    if req.Theme == "" {
        // Return error for empty theme
        return nil, status.Errorf(codes.InvalidArgument, "theme cannot be empty")
    }

    // TODO: Add actual logic to switch theme
    fmt.Printf("Theme changed to: %s
", req.Theme)

    // Return success response
    return &ThemeResponse{Success: true}, nil
}

// RegisterThemeService creates and registers the ThemeService server.
func RegisterThemeService(lis net.Listener) {
    // Create a new server instance
    server := grpc.NewServer()

    // Register the service
    RegisterThemeServiceServer(server, &ThemeServer{})

    // Start serving.
    if err := server.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Register the service and start serving
    RegisterThemeService(lis)
}
