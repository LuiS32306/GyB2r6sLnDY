// 代码生成时间: 2025-10-01 03:14:27
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/timestamppb"
)

// LeaderboardService is the server implementation of leaderboard service.
type LeaderboardService struct {
    // server stream is used for server-side streaming
    serverStream grpc.ServerStream
}

// Leaderboard defines the protobuf message for leaderboard entries.
type Leaderboard struct {
    Id      int64                        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
    Name    string                      `protobuf:"string,2,opt,name=name,proto3" json:"name,omitempty"`
    Score   int32                        `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
    Updated *timestamppb.Timestamp      `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
}

// LeaderboardResponse defines the protobuf message for leaderboard response.
type LeaderboardResponse struct {
    Entries []*Leaderboard `protobuf:"repeated,1,opt,name=entries,proto3" json:"entries,omitempty"`
}

// RegisterServer registers the leaderboard service to the gRPC server.
func RegisterServer(s *grpc.Server, server *LeaderboardService) {
    leaderboardpb.RegisterLeaderboardServer(s, server)
}

// GetAllEntries is a gRPC method to get all leaderboard entries.
func (s *LeaderboardService) GetAllEntries(ctx context.Context, in *leaderboardpb.GetAllEntriesRequest) (*LeaderboardResponse, error) {
    // Simulating database calls with a map for simplicity.
    // In production, replace with actual database calls.
    var entries []*Leaderboard
    // Assuming we have a function to fetch leaderboard entries.
    entries = FetchLeaderboardEntries()
    
    return &LeaderboardResponse{Entries: entries}, nil
}

// FetchLeaderboardEntries simulates fetching entries from a database.
// In a real-world scenario, this would involve database queries.
func FetchLeaderboardEntries() []*Leaderboard {
    // Sample data for demonstration purposes.
    return []*Leaderboard{
        {Id: 1, Name: "Alice", Score: 100},
        {Id: 2, Name: "Bob", Score: 200},
        {Id: 3, Name: "Charlie", Score: 300},
    }
}

// RunServer starts the gRPC server and blocks until it is terminated.
func RunServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    s := grpc.NewServer()
    RegisterServer(s, &LeaderboardService{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

func main() {
    // Wait for interrupt signal to gracefully shutdown the server.
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        // Here you might want to add logic to gracefully shutdown
        // any running goroutines, close database connections, etc.
        fmt.Println("Server is shutting down...")
        // Perform any necessary cleanup here.
    }()

    RunServer()
}
