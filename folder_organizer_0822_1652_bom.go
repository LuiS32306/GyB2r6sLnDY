// 代码生成时间: 2025-08-22 16:52:46
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "path/filepath"
    "sort"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "google.golang.org/protobuf/types/known/timestamppb"
)

// Define the constants for the service
const (
    port = ":50051"
)

// FolderOrganizerService provides methods to organize folder structures.
type FolderOrganizerService struct{}

// OrganizeFolder method organizes a folder by sorting files and directories.
func (s *FolderOrganizerService) OrganizeFolder(ctx context.Context, req *OrganizeRequest) (*OrganizeResponse, error) {
    start := time.Now()
    defer func() {
        fmt.Printf("OrganizeFolder took %v", time.Since(start))
    }()

    // Check if the directory exists.
    if _, err := os.Stat(req.Directory); os.IsNotExist(err) {
        return nil, fmt.Errorf("directory does not exist: %w", err)
    }

    // Read the directory contents.
    entries, err := os.ReadDir(req.Directory)
    if err != nil {
        return nil, fmt.Errorf("failed to read directory: %w", err)
    }

    // Sort the directory contents.
    var files []os.DirEntry
    for _, entry := range entries {
        files = append(files, entry)
    }
    sort.Slice(files, func(i, j int) bool {
        return files[i].Name() < files[j].Name()
    })

    // Create a new directory for organized files if needed.
    if req.OutputDirectory != "" {
        if err := os.MkdirAll(req.OutputDirectory, 0755); err != nil {
            return nil, fmt.Errorf("failed to create output directory: %w", err)
        }
    }

    // Move files to the output directory.
    for _, file := range files {
        sourcePath := filepath.Join(req.Directory, file.Name())
        destPath := filepath.Join(req.OutputDirectory, file.Name())
        if err := os.Rename(sourcePath, destPath); err != nil {
            return nil, fmt.Errorf("failed to move file: %w", err)
        }
    }

    return &OrganizeResponse{
        Message:   "Folder organized successfully.",
        Elapsed:   time.Since(start).Seconds(),
       .Timestamp: timestamppb.Now(),
    }, nil
}

// Main function to start the GRPC server.
func main() {
    fmt.Println("Starting Folder Organizer GRPC Service...")
    lis, err := net.Listen("tcp\, port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port", port)

    // Create a new GRPC server.
    s := grpc.NewServer()

    // Register the service with the server.
    RegisterFolderOrganizerServiceServer(s, &FolderOrganizerService{})

    // Register reflection service on gRPC server.
    reflection.Register(s)

    // Start the GRPC server.
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// OrganizeRequest is the request message for the OrganizeFolder method.
type OrganizeRequest struct {
    Directory      string
    OutputDirectory string
}

// OrganizeResponse is the response message for the OrganizeFolder method.
type OrganizeResponse struct {
    Message   string
    Elapsed  float64
    Timestamp *timestamppb.Timestamp
}

// RegisterFolderOrganizerServiceServer registers the FolderOrganizerService with the GRPC server.
func RegisterFolderOrganizerServiceServer(s *grpc.Server, srv *FolderOrganizerService) {
    RegisterFolderOrganizerServiceServer(s, srv)
}
