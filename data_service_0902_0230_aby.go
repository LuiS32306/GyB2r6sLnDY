// 代码生成时间: 2025-09-02 02:30:45
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

// DataModel is the data model that we'll be using for this example.
type DataModel struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// DataServer is the server API for DataService service.
type DataServer struct {
    // Embed nil to allow DataServer to implement DataServiceServer.
    grpc.UnimplementedDataServiceServer
}

// NewDataServer creates a new instance of DataServer.
func NewDataServer() *DataServer {
    return &DataServer{}
}

// GetData is a method to get a DataModel by its ID.
func (s *DataServer) GetData(ctx context.Context, req *GetRequest) (*DataModel, error) {
    // Simulate a database call
    dataModel := DataModel{ID: req.Id, Name: "Example Name"}

    // Return the data model
    return &dataModel, nil
}

// GetRequest is the request message for the Data service.
type GetRequest struct {
    Id int `json:"id"`
}

// DataService provides operations for data handling.
type DataService interface {
    GetData(ctx context.Context, req *GetRequest) (*DataModel, error)
}

func main() {
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")

    // Create a new gRPC server
    srv := grpc.NewServer()

    // Register the data service on the server
    dataServer := NewDataServer()
    RegisterDataServiceServer(srv, dataServer)

    // Start the server
    if err := srv.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// RegisterDataServiceServer registers the DataServiceServer to the gRPC server.
func RegisterDataServiceServer(srv *grpc.Server, service DataService) {
    RegisterDataServiceServer(srv, service)
}

// The following code is the service definition in a .proto file which is needed to generate the gRPC code.
// // Data service definition
// service DataService {
//     // Gets a data model by ID.
//     rpc GetData(GetRequest) returns (DataModel) {}
// }
// 
// // The request message containing the user's ID.
// message GetRequest {
//     // Unique identifier for the data model.
//     int32 id = 1;
// // }
// 
// // The response message containing the data model.
// message DataModel {
//     // Unique identifier for the data model.
//     int32 id = 1;
//     // The user's name.
//     string name = 2;
// // }