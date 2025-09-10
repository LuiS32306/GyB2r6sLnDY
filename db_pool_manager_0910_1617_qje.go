// 代码生成时间: 2025-09-10 16:17:36
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

// DatabaseConfig contains the database connection parameters
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// DBPool is a struct that wraps the *sql.DB connection pool
type DBPool struct {
    *sql.DB
    cfg DatabaseConfig
}

// NewDBPool creates a new database connection pool
func NewDBPool(cfg DatabaseConfig) (*DBPool, error) {
    // Construct the DSN (Data Source Name) for MySQL
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

    // Open the database connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Set connection pool parameters
    db.SetMaxIdleConns(10) // Maximum number of connections in the idle connection pool.
    db.SetMaxOpenConns(100) // Maximum number of open connections to the database.
    db.SetConnMaxLifetime(30 * time.Minute) // Maximum amount of time a connection may be reused.

    // Ping the database to verify the connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return &DBPool{DB: db, cfg: cfg}, nil
}

func main() {
    // Define the database connection configuration
    dbConfig := DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "user",
        Password: "password",
        DBName:   "dbname",
    }

    // Create a new database connection pool
    dbPool, err := NewDBPool(dbConfig)
    if err != nil {
        log.Fatalf("Failed to create database pool: %s", err)
    }
    defer dbPool.Close() // Ensure the pool is closed before exiting the program

    // Example usage of the database pool
    fmt.Println("Database connection pool created successfully.")
    // Here you would add your business logic using dbPool.DB to execute queries
}
