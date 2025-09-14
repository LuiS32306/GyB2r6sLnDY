// 代码生成时间: 2025-09-14 15:58:37
package main

import (
    "context"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
    "time"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
)

// AutomatedTestSuite provides a struct to manage the gRPC server and client interactions.
type AutomatedTestSuite struct {
    Server *grpc.Server
    Client AutomatedTestServiceClient
    Conn   *grpc.ClientConn
}

// AutomatedTestServiceClient is the client API for AutomatedTestService service.
type AutomatedTestServiceClient interface {
    ExecuteTest(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

// TestRequest is the request message to execute a test.
type TestRequest struct {
    TestName string
    Params   map[string]string
}

// TestResponse is the response message from executing a test.
type TestResponse struct {
    Result  string
    Success bool
}

// StartServer starts the gRPC server.
func (ats *AutomatedTestSuite) StartServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    ats.Server = grpc.NewServer()
    RegisterAutomatedTestServiceServer(ats.Server, &server{})
    if err := ats.Server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// StopServer stops the gRPC server.
func (ats *AutomatedTestSuite) StopServer() {
    ats.Server.Stop()
}

// ConnectToServer connects to the gRPC server.
func (ats *AutomatedTestSuite) ConnectToServer(address string) error {
    var err error
    ats.Conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        return err
    }
    ats.Client = NewAutomatedTestServiceClient(ats.Conn)
    return nil
}

// DisconnectFromServer disconnects from the gRPC server.
func (ats *AutomatedTestSuite) DisconnectFromServer() {
    ats.Conn.Close()
}

// RunTest runs a test using the client.
func (ats *AutomatedTestSuite) RunTest(testName string, params map[string]string) (*TestResponse, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    request := &TestRequest{TestName: testName, Params: params}
    response, err := ats.Client.ExecuteTest(ctx, request)
    return response, err
}

// NewAutomatedTestSuite creates a new instance of AutomatedTestSuite.
func NewAutomatedTestSuite() *AutomatedTestSuite {
    return &AutomatedTestSuite{}
}

// server is used to implement AutomatedTestServiceServer.
type server struct{}

// ExecuteTest is a method of AutomatedTestServiceServer.
func (s *server) ExecuteTest(ctx context.Context, request *TestRequest) (*TestResponse, error) {
    // Implement your test execution logic here. For now, return a mock response.
    return &TestResponse{Result: "Test executed", Success: true}, nil
}

// RegisterAutomatedTestServiceServer registers AutomatedTestServiceServer to gRPC server.
func RegisterAutomatedTestServiceServer(s *grpc.Server, srv AutomatedTestServiceServer) {
    RegisterAutomatedTestServiceHandler(s, srv)
}

func main() {
    ats := NewAutomatedTestSuite()
    go ats.StartServer()
    defer ats.StopServer()

    if err := ats.ConnectToServer("localhost:50051"); err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer ats.DisconnectFromServer()

    testResponse, err := ats.RunTest("SampleTest", nil)
    if err != nil {
        log.Fatalf("could not run test: %v", err)
    }
    log.Printf("Test Result: %s, Success: %t", testResponse.Result, testResponse.Success)

    // Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    <-ch
    log.Println("Shutting down server...")
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    ats.StopServer()
    <-shutdownCtx.Done()
    log.Println("Server exiting")
}