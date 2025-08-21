// 代码生成时间: 2025-08-22 04:29:45
package main

import (
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/protobuf/types/known/anypb"
    "google.golang.org/protobuf/types/known/emptypb"
)

// User represents a user data model.
type User struct {
    Id      string
    Name    string
    Email   string
    Created int64
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer struct {
    // Define any additional fields here
}

// NewUserServiceServer creates a new instance of the service.
func NewUserServiceServer() *UserServiceServer {
    return &UserServiceServer{}
}

// CreateUser creates a new user.
func (s *UserServiceServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
    if req == nil || req.User == nil {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request")
    }
    // Add your logic to create a user here
    // For simplicity, just return the received user
    return req.User, nil
}

// GetUser retrieves a user by ID.
func (s *UserServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
    if req == nil || req.Id == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request\)
    }
    // Add your logic to retrieve a user by ID here
    // For simplicity, just return a dummy user
    return &User{Id: req.Id}, nil
}

// UpdateUser updates an existing user.
func (s *UserServiceServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
    if req == nil || req.User == nil {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request")
    }
    // Add your logic to update a user here
    // For simplicity, just return the received user
    return req.User, nil
}

// DeleteUser deletes a user by ID.
func (s *UserServiceServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*emptypb.Empty, error) {
    if req == nil || req.Id == "" {
        return nil, status.Errorf(codes.InvalidArgument, "invalid request")
    }
    // Add your logic to delete a user by ID here
    return &emptypb.Empty{}, nil
}

// RegisterUserServer registers this server with the GRPC server.
func RegisterUserServer(server *grpc.Server) {
    RegisterUserServiceServer(server, NewUserServiceServer())
}

// Below are the gRPC request and response types for the User service.

// CreateUserRequest is the request for creating a user.
type CreateUserRequest struct {
    User *User `protobuf:"bytes,1,opt,name=user,proto3"`
}

// GetUserRequest is the request for getting a user.
type GetUserRequest struct {
    Id string `protobuf:"bytes,1,opt,name=id,proto3"`
}

// UpdateUserRequest is the request for updating a user.
type UpdateUserRequest struct {
    User *User `protobuf:"bytes,1,opt,name=user,proto3"`
}

// DeleteUserRequest is the request for deleting a user.
type DeleteUserRequest struct {
    Id string `protobuf:"bytes,1,opt,name=id,proto3"`
}

// The following comments are the protocol buffer definitions for the service and messages.
// Please replace them with the actual generated code from the .proto file.

// message User {
//     string id = 1;
//     string name = 2;
//     string email = 3;
//     int64 created = 4;
// }

// message CreateUserRequest {
//     User user = 1;
// }

// message GetUserRequest {
//     string id = 1;
// }

// message UpdateUserRequest {
//     User user = 1;
// }

// message DeleteUserRequest {
//     string id = 1;
// }

func main() {
    fmt.Println("Starting user service...")
    // Define the service server and register it with the GRPC server
    server := grpc.NewServer()
    RegisterUserServer(server)
    // Define the service address and start the server
    fmt.Println("Serving on port 50051")
    if err := server.Serve(lis); err != nil {
        fmt.Println("Failed to serve: ", err)
    }
}
