// 代码生成时间: 2025-09-24 13:13:07
// log_parser.go - A GRPC service for parsing log files.
    
package main


import (
    "context"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// LogLine represents a single line from a log file.
type LogLine struct {
    Timestamp string
    Level     string
    Message   string
}

// LogFile represents a log file that needs to be parsed.
type LogFile struct {
    Path string
}

// LogParserServer defines the server methods for the LogParser service.
type LogParserServer struct{}

// ParseLogFile is called by clients to parse a log file.
func (s *LogParserServer) ParseLogFile(ctx context.Context, req *LogFile) (*[]LogLine, error) {
    // Read the log file.
    data, err := ioutil.ReadFile(req.Path)
    if err != nil {
        return nil, status.Errorf(codes.InvalidArgument, "Unable to read file: %v", err)
    }

    // Split the file into lines.
    lines := strings.Split(strings.TrimSpace(string(data)), "\
")

    // Initialize a slice to hold the parsed lines.
    var parsedLines []LogLine

    // Iterate over each line and parse it.
    for _, line := range lines {
        parts := strings.Fields(line)
        if len(parts) < 3 {
            // Skip lines that don't have enough parts.
            continue
       }

        // Assuming the first part is the timestamp, the second is the log level, and the rest is the message.
        parsedLines = append(parsedLines, LogLine{
            Timestamp: parts[0],
            Level:     parts[1],
            Message:   strings.Join(parts[2:], " "),
        })
    }

    return &parsedLines, nil
}

func main() {
    // Define the port on which the server will run.
    port := ":50051"

    // Create a new server.
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("Failed to listen on port %s: %v", port, err)
    }

    // Create a new gRPC server.
    grpcServer := grpc.NewServer()

    // Register the LogParserServer service.
    LogParserServer := &LogParserServer{}
    pb.RegisterLogParserServer(grpcServer, LogParserServer)

    // Start the server.
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
