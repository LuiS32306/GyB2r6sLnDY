// 代码生成时间: 2025-07-31 23:08:22
package main

import (
    "context"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "testing"
    "time"

    "google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
    // helloworld.UnimplementedGreeterServer is embedded.
    helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
    // Your logic here.
    return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// TestServer is a test for the grpc server.
func TestServer(t *testing.T) {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    defer lis.Close()

    s := grpc.NewServer()
    defer s.Stop()

    helloworld.RegisterGreeterServer(s, &server{})

    go func() {
        if err := s.Serve(lis); err != nil {
            log.Fatalf("failed to serve: %v", err)
        }
    }()

    // Set a timeout for the test.
    timeout := time.After(10 * time.Second)
    for {
        select {
        case <-timeout:
            t.Fatalf("timeout while trying to connect to server")
        default:
            // Try to connect to the server.
            conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
            if err == nil {
                conn.Close()
                return
            }
            time.Sleep(100 * time.Millisecond)
        }
    }
}

func main() {
    // Run the test.
    testing.Main(
        func(pat, str string) bool { return true },  // Match all files.
        func(*testing.M) {},                      // No setup.
        func() {
            TestServer(nil)  // Run the test.
        },
    )
}
