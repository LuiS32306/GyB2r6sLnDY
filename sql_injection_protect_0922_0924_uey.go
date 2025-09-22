// 代码生成时间: 2025-09-22 09:24:26
package main

import (
    "database/sql"
    "fmt"
    "log"
    "strings"
    "time"
    "github.com/go-sql-driver/mysql"
)

// User represents a user entity with fields that match the database schema.
type User struct {
    ID       uint      `db:"id"`
    Username string    `db:"username"`
    Email    string    `db:"email"`
    CreatedAt time.Time `db:"created_at"`
}

// NewMySQLDriver returns a new MySQL driver instance for sql.DB.
func NewMySQLDriver() *sql.DB {
    // Connect to the database using the MySQL driver.
    // Replace the placeholder values with your actual database credentials.
    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatalf("Error connecting to the database: %s
", err)
    }
    // Set the maximum number of connections in the idle connection pool.
    db.SetMaxIdleConns(10)
    return db
}

// GetUserByUsername queries the database for a user by username, preventing SQL injection.
func GetUserByUsername(db *sql.DB, username string) (*User, error) {
    // Prepare a SQL statement to prevent SQL injection.
    stmt, err := db.Prepare("SELECT * FROM users WHERE username = ?")
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    // Execute the prepared statement with the username parameter.
    var user User
    err = stmt.QueryRow(username).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil // No user found with the given username.
        }
        return nil, err
    }
    return &user, nil
}

// main function to demonstrate the prevention of SQL injection.
func main() {
    db := NewMySQLDriver()
    defer db.Close()

    // Example usage: retrieve a user by their username.
    user, err := GetUserByUsername(db, "exampleUser")
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("User not found.")
        } else {
            fmt.Printf("An error occurred: %s
", err)
       }
    } else {
        fmt.Printf("User found: %+v
", user)
    }
}
