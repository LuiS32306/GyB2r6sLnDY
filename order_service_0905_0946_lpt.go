// 代码生成时间: 2025-09-05 09:46:10
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

// Define the OrderServiceServer which will implement the OrderServiceServer interface
type OrderServiceServer struct {}

// ProcessOrder is a method of the OrderServiceServer interface
func (s *OrderServiceServer) ProcessOrder(ctx context.Context, in *OrderRequest) (*emptypb.Empty, error) {
    // Check if order details are provided
    if in.GetOrderId() == 0 || in.GetCustomerId() == 0 {
        return nil, status.Errorf(codes.InvalidArgument, "Order ID and Customer ID must be provided")
    }

    // Simulate order processing logic here
    // For demonstration, just log the order details and return success
    log.Printf("Processing order ID: %d for customer ID: %d", in.GetOrderId(), in.GetCustomerId())

    // Return an empty response to indicate success
    return &emptypb.Empty{}, nil
}

// StartServer starts the GRPC server
func StartServer(lis net.Listener, done chan bool) {
    // Create a new GRPC server
    srv := grpc.NewServer()

    // Register the OrderServiceServer to the server
    RegisterOrderServiceServer(srv, &OrderServiceServer{})

    // Start serving
    if err := srv.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}

func main() {
    // Create a listener on TCP port 50051
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    // Create a channel to handle server shutdown
    done := make(chan bool, 1)

    // Start the GRPC server
    go StartServer(lis, done)

    // Wait for the server to be closed
    <-done
}

// OrderRequest message definition
type OrderRequest struct {
    OrderId   int64  `protobuf:"varint,1,opt,name=order_id,json=orderId"`
    CustomerId int64 `protobuf:"varint,2,opt,name=customer_id,json=customerId"`
}

// OrderServiceServer interface definition
type OrderServiceServer interface {
    ProcessOrder(context.Context, *OrderRequest) (*emptypb.Empty, error)
}

// RegisterOrderServiceServer registers the OrderServiceServer to the GRPC server
func RegisterOrderServiceServer(srv *grpc.Server, srvImpl OrderServiceServer) {
    srv.RegisterService(&_OrderService_serviceDesc, srvImpl)
}

// _OrderService_serviceDesc is the descriptor for the OrderService service
var _OrderService_serviceDesc = grpc.ServiceDesc{
    ServiceName: "OrderService",
    HandlerType: (*OrderServiceServer)(nil),
    Methods: []grpc.MethodDesc{
        {
            MethodName: "ProcessOrder",
            Handler: func (srv interface{}, ctx context.Context,
                in1, in2 interface{}) (interface{}, error) {
                return srv.(OrderServiceServer).ProcessOrder(ctx, in1.(*OrderRequest))
            },
        },
    },
    Streams:  []grpc.StreamDesc{},
    Metadata: "order_service.proto",
}
