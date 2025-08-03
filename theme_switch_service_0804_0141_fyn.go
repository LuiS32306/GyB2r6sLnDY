// 代码生成时间: 2025-08-04 01:41:43
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/proto"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// Theme defines the type of theme that can be used
type Theme int

const (
    // DarkTheme represents a dark theme
    DarkTheme Theme = iota
    // LightTheme represents a light theme
    LightTheme
)

// ThemeService provides methods for switching themes
type ThemeService struct {
    // currentTheme stores the currently selected theme
    currentTheme Theme
}

// NewThemeService creates a new ThemeService with the default theme
func NewThemeService() *ThemeService {
    return &ThemeService{currentTheme: LightTheme}
}

// SwitchTheme switches the current theme to the opposite of what it currently is
func (s *ThemeService) SwitchTheme(ctx context.Context, in *SwitchThemeRequest) (*SwitchThemeResponse, error) {
    // Check if the request is valid
    if in == nil {
        return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
    }

    // Determine the new theme based on the current theme
    newTheme := DarkTheme
    if s.currentTheme == DarkTheme {
        newTheme = LightTheme
    }

    // Update the current theme
    s.currentTheme = newTheme

    // Return the new theme as the response
    return &SwitchThemeResponse{Theme: newTheme}, nil
}

// SwitchThemeRequest is the request message for the SwitchTheme method
type SwitchThemeRequest struct{}

// SwitchThemeResponse is the response message for the SwitchTheme method
type SwitchThemeResponse struct {
    // Theme is the new theme after switching
    Theme Theme
}

// RegisterThemeServiceServer registers the ThemeService with the gRPC server
func RegisterThemeServiceServer(s *grpc.Server, service *ThemeService) {
    // Register the service with the gRPC server
    themeServiceServer := &themeServiceServer{service}
    RegisterThemeServiceServer(s, themeServiceServer)
}

type themeServiceServer struct {
    ThemeService
}

// SwitchTheme implements ThemeServiceServer
func (s *themeServiceServer) SwitchTheme(ctx context.Context, in *SwitchThemeRequest) (*SwitchThemeResponse, error) {
    return s.ThemeService.SwitchTheme(ctx, in)
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on :50051")

    server := grpc.NewServer()
    defer server.Stop()

    // Create a new theme service
    themeService := NewThemeService()

    // Register the theme service with the gRPC server
    RegisterThemeServiceServer(server, themeService)

    // Start the server
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}