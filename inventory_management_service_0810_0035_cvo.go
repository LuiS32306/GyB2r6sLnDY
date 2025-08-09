// 代码生成时间: 2025-08-10 00:35:36
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

// InventoryItem represents an item in the inventory
type InventoryItem struct {
    ID          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
    Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
    Quantity   int32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
    Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
    SKU         string  `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
    Price       float32 `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
}

// InventoryService provides operations for managing inventory items
type InventoryService struct {
    // In a real-world application, this would be a database connection or an in-memory store
    inventory map[string]InventoryItem
}

// AddItem adds a new item to the inventory
func (s *InventoryService) AddItem(ctx context.Context, item *InventoryItem) (*InventoryItem, error) {
    if item.ID == "" {
        return nil, status.Errorf(codes.InvalidArgument, "item ID cannot be empty")
    }
    s.inventory[item.ID] = *item
    return item, nil
}

// GetItem retrieves an item from the inventory by its ID
func (s *InventoryService) GetItem(ctx context.Context, id *emptypb.Empty) (*InventoryItem, error) {
    // In a real-world application, you would query the database or in-memory store
    // For demonstration purposes, we assume the ID is the key
    for _, item := range s.inventory {
        return &item, nil
    }
    return nil, status.Errorf(codes.NotFound, "item not found")
}

// UpdateItem updates an existing item in the inventory
func (s *InventoryService) UpdateItem(ctx context.Context, item *InventoryItem) (*InventoryItem, error) {
    if item.ID == "" {
        return nil, status.Errorf(codes.InvalidArgument, "item ID cannot be empty")
    }
    if _, exists := s.inventory[item.ID]; !exists {
        return nil, status.Errorf(codes.NotFound, "item not found")
    }
    s.inventory[item.ID] = *item
    return item, nil
}

// DeleteItem removes an item from the inventory by its ID
func (s *InventoryService) DeleteItem(ctx context.Context, id *emptypb.Empty) (*emptypb.Empty, error) {
    // In a real-world application, you would query the database or in-memory store
    // For demonstration purposes, we assume the ID is the key
    for key := range s.inventory {
        delete(s.inventory, key)
        return &emptypb.Empty{}, nil
    }
    return nil, status.Errorf(codes.NotFound, "item not found")
}

// server is used to implement inventory.InventoryServiceServer
type server struct {
    inventory.UnimplementedInventoryServiceServer
}

// NewServer creates a new instance of the server
func NewServer() *server {
    return &server{}
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Server listening on port 50051")
    
    s := grpc.NewServer()
    inventory.RegisterInventoryServiceServer(s, NewServer())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
