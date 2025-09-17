// 代码生成时间: 2025-09-17 19:55:33
package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/reflection"
)

// ErrorLog defines the service for collecting error logs
type ErrorLog struct{}

// CollectErrorLog is a method for collecting error logs
func (e *ErrorLog) CollectErrorLog(ctx context.Context, log *ErrorLogRequest) (*Empty, error) {
    // Check if the log request is nil
    if log == nil {
        return nil, status.Errorf(codes.InvalidArgument, "empty error log request")
    }

    // Log the error message to standard output
    fmt.Printf("Error: %s
", log.Message)

    // Write the error log to a file for persistence
    err := writeLogToFile(log.Message)
    if err != nil {
        return nil, status.Errorf(codes.Internal, "failed to write log to file: %v", err)
    }

    // Return an Empty message as we don't need to send back any data
    return &Empty{}, nil
}

// writeLogToFile writes the message to a log file
func writeLogToFile(message string) error {
    // Open the log file
    file, err := os.OpenFile("error_logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    // Write the time and message to the file
    if _, err := file.WriteString(time.Now().Format("") + " - " + message + "
"); err != nil {
        return err
    }

    return nil
}

// Empty is a placeholder message
type Empty struct{}

// ErrorLogRequest is a message that contains the error log
type ErrorLogRequest struct {
    Message string `protobuf:"bytes,1,opt,name=message"`
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")

    s := grpc.NewServer()
    // Register the service
    RegisterErrorLogServiceServer(s, &ErrorLog{})
    reflection.Register(s)

    // Trap SIGINT to gracefully shutdown the server
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("")
        fmt.Println("Shutting down server...")
        s.GracefulStop()
    }()

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
