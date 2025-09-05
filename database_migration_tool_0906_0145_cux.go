// 代码生成时间: 2025-09-06 01:45:14
package main

import (
    "context"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "google.golang.org/grpc"
    "io"
    "log"
    "os"
)

// DatabaseMigrationService provides methods for database migration
type DatabaseMigrationService struct {
    db *sql.DB
}

// NewDatabaseMigrationService creates a new instance of DatabaseMigrationService
func NewDatabaseMigrationService(db *sql.DB) *DatabaseMigrationService {
    return &DatabaseMigrationService{db: db}
}

// MigrateDatabase performs database migration
func (s *DatabaseMigrationService) MigrateDatabase(ctx context.Context, in *MigrationRequest) (*MigrationResponse, error) {
    if in == nil {
        return nil, fmt.Errorf("migration request cannot be nil")
    }
    
    // Execute migration SQL commands
    _, err := s.db.ExecContext(ctx, in.Sql)
    if err != nil {
        return nil, err
    }
    
    // Return success response
    return &MigrationResponse{Success: true}, nil
}

// MigrationRequest is the request message for the MigrateDatabase method
type MigrationRequest struct {
    Sql string
}

// MigrationResponse is the response message for the MigrateDatabase method
type MigrationResponse struct {
    Success bool
}

// main function
func main() {
    // Set up database connection
    db, err := sql.Open("mysql", "user:password@tcp(host:port)/dbname")
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer db.Close()
    
    // Create GRPC server
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")
    
    // Create database migration service
    migrationService := NewDatabaseMigrationService(db)
    
    // Start GRPC server
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// GRPC Server setup
func startGRPCServer(migrationService *DatabaseMigrationService) *grpc.Server {
    s := grpc.NewServer()
    
    // Register MigrationService on the server
    migratepb.RegisterMigrationServiceServer(s, migrationService)
    
    return s
}

// Below is the proto file for the migration service
// Assuming the proto file is named migration.proto
//
// syntax = "proto3";
//
// package migrate;
//
// service MigrationService {
//     rpc MigrateDatabase(MigrationRequest) returns (MigrationResponse) {}
// }
//
// message MigrationRequest {
//     string sql = 1;
// }
//
// message MigrationResponse {
//     bool success = 1;
// }
