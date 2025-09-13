// 代码生成时间: 2025-09-14 06:49:48
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// TestData represents the structure of the test data
type TestData struct {
    ID       int
    Name     string
    Email    string
    Password string
}

// RandomString generates a random string with a specified length
func RandomString(length int) string {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

// GenerateTestData creates a TestData with random values
func GenerateTestData() TestData {
    rand.Seed(time.Now().UnixNano()) // Seed the random number generator
    return TestData{
        ID:       rand.Int(),
        Name:     RandomString(10), // Generate a random string of length 10 for the name
        Email:    fmt.Sprintf("%s@example.com", RandomString(10)), // Generate a random email
        Password: RandomString(12), // Generate a random string of length 12 for the password
    }
}

func main() {
    testData := GenerateTestData()
    fmt.Printf("Generated Test Data: %+v
", testData)
}
