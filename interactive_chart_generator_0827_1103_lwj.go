// 代码生成时间: 2025-08-27 11:03:40
package main

import (
    "context"
    "fmt"
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/emptypb"
)

// ChartInfo contains the necessary information to generate a chart
type ChartInfo struct {
    Title   string   `protobuf:"bytes,1,opt,name=title,proto3"`
    XLabels []string `protobuf:"bytes,2,rep,name=xLabels,proto3"`
    YValues []int32  `protobuf:"varint,3,rep,packed,name=yValues,proto3"`
}

// ChartServiceServer is the server API for ChartService service
type ChartServiceServer struct {
    // Embed	emptypb.Empty to allow for embedding
    emptypb.Empty_unverified
}

// GenerateChart generates an interactive chart based on the provided ChartInfo
func (s *ChartServiceServer) GenerateChart(ctx context.Context, in *ChartInfo) (*emptypb.Empty, error) {
    // Check if the input ChartInfo is valid
    if in.Title == "" || len(in.XLabels) == 0 || len(in.YValues) == 0 {
        return nil, fmt.Errorf("invalid chart information")
    }

    // Logic to generate the chart would go here
    // For demonstration purposes, we just print the chart information to the console
    fmt.Printf("Generating chart with title: %s\
", in.Title)
    fmt.Printf("X Labels: %v\
", in.XLabels)
    fmt.Printf("Y Values: %v\
", in.YValues)

    // Return an empty response to indicate success
    return &emptypb.Empty{}, nil
}

// main function to start the gRPC server
func main() {
   lis, err := net.Listen("tcp", ":50051")
   if err != nil {
       log.Fatalf("failed to listen: %v", err)
   }

   // Create a new gRPC server
   srv := grpc.NewServer()

   // Register the ChartServiceServer
   // Assuming ChartServiceServer is registered with the grpc codegen tool
   RegisterChartServiceServer(srv, &ChartServiceServer{})

   // Start the gRPC server
   if err := srv.Serve(lis); err != nil {
       log.Fatalf("failed to serve: %v", err)
   }
}

// RegisterChartServiceServer registers the ChartServiceServer to the provided gRPC Server
func RegisterChartServiceServer(srv *grpc.Server, srvImpl ChartServiceServer) {
    // Register the service to the server
    RegisterChartServiceServer(srv, srvImpl)
}

// Note: This code assumes the existence of a ChartService proto file,
// which defines the service interface and the ChartInfo message.
// The codegen tool generates the necessary protobuf code for golang.

// The actual implementation of chart generation would depend on the specific
// libraries and frameworks used for creating interactive charts.
