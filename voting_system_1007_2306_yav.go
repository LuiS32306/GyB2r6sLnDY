// 代码生成时间: 2025-10-07 23:06:50
@author: Your Name
@email: your.email@example.com
*/

package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

// Define the voting service
type VotingServiceServer struct {
    // Store the vote count for each option
    votes map[string]int
}

// Define the service methods
type votingServiceServer struct{
    // UnimplementedVotingServiceServer must be embedded to have forward compatible methods
    UnimplementedVotingServiceServer
}

// Register the service with gRPC
func (s *votingServiceServer) CastVote(ctx context.Context, in *CastVoteRequest) (*emptypb.Empty, error) {
    option := in.getOption()
    if _, exists := s.votes[option]; !exists {
        s.votes[option] = 0
    }
    s.votes[option]++
    return &emptypb.Empty{}, nil
}

func (s *votingServiceServer) GetResults(ctx context.Context, in *emptypb.Empty) (*GetResultsResponse, error) {
    results := make(map[string]int)
    for option, count := range s.votes {
        results[option] = count
    }
    return &GetResultsResponse{Results: results}, nil
}

// Initialize the server
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server is running on port 50051")

    s := grpc.NewServer()
    // Register the voting service with the server
    RegisterVotingServiceServer(s, &votingServiceServer{votes: make(map[string]int)})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Define the protobuf messages and services in a separate file
// voting_service.proto
// message CastVoteRequest {
//     string option = 1;
// }
// message GetResultsResponse {
//     map<string, int32> results = 1;
// }
// service VotingService {
//     rpc CastVote(CastVoteRequest) returns (google.protobuf.Empty);
//     rpc GetResults(google.protobuf.Empty) returns (GetResultsResponse);
// }