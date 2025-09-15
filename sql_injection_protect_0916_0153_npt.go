// 代码生成时间: 2025-09-16 01:53:59
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "strings"
    "time"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

// User represents a user entity with fields that can be used for database queries.
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

// Database represents a database connection with a predefined SQL dialect.
type Database struct {
    db *sql.DB
}

// NewDatabase initializes and returns a new Database instance.
func NewDatabase(dataSourceName string) (*Database, error) {
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, err
    }
    return &Database{db: db}, nil
}

// QueryUser retrieves a user by their ID, ensuring SQL injection prevention.
func (d *Database) QueryUser(ctx context.Context, userID int) (*User, error) {
    // Prepare the SQL statement to prevent SQL injection by using placeholders.
    stmt, err := d.db.PrepareContext(ctx, "SELECT id, username, email FROM users WHERE id = ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()
    
    var user User
    err = stmt.QueryRowContext(ctx, userID).Scan(&user.ID, &user.Username, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // No user found, but not an error.
        }
        return nil, err
    }
    return &user, nil
}

// InsertUser inserts a new user into the database, avoiding SQL injection.
func (d *Database) InsertUser(ctx context.Context, user *User) error {
    // Prepare the SQL statement to avoid SQL injection.
    stmt, err := d.db.PrepareContext(ctx, "INSERT INTO users (username, email) VALUES (?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()
    
    _, err = stmt.ExecContext(ctx, user.Username, user.Email)
    return err
}

func main() {
    // Example usage of the database operations to prevent SQL injection.
    dataSourceName := "user:password@tcp(127.0.0.1:3306)/dbname"
    db, err := NewDatabase(dataSourceName)
    if err != nil {
        log.Fatal(err)
    }
    defer db.db.Close()
    
    // Set up a context with a timeout for database operations.
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    // Query a user by ID.
    user, err := db.QueryUser(ctx, 1)
    if err != nil {
        log.Printf("Error querying user: %v", err)
    } else if user != nil {
        fmt.Printf("Found user: %+v
", user)
    }
    
    // Insert a new user.
    newUser := User{Username: "john", Email: "john@example.com"}
    if err := db.InsertUser(ctx, &newUser); err != nil {
        log.Printf("Error inserting user: %v", err)
    } else {
        fmt.Println("User inserted successfully")
    }
}