// 代码生成时间: 2025-10-11 01:53:25
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// MockDataGenerator is the main structure for generating mock data.
type MockDataGenerator struct {}

// NewMockDataGenerator creates a new instance of MockDataGenerator.
func NewMockDataGenerator() *MockDataGenerator {
    return &MockDataGenerator{}
}

// GenerateRandomString generates a random string of a specified length.
func (g *MockDataGenerator) GenerateRandomString(length int) string {
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    randomString := make([]rune, length)
    for i := range randomString {
        randomString[i] = letters[rand.Intn(len(letters))]
    }
    return string(randomString)
}

// GenerateRandomInt generates a random integer within a specified range.
func (g *MockDataGenerator) GenerateRandomInt(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
}

// GenerateRandomBool generates a random boolean value.
func (g *MockDataGenerator) GenerateRandomBool() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) == 0
}

// GenerateUserData generates mock user data.
func (g *MockDataGenerator) GenerateUserData() *UserData {
    return &UserData{
        Username: g.GenerateRandomString(10),
        Email:    g.GenerateRandomString(20) + "@example.com",
        Age:      g.GenerateRandomInt(18, 65),
        IsAdmin:  g.GenerateRandomBool(),
    }
}

// UserData represents the structure of user data.
type UserData struct {
    Username string
    Email    string
    Age      int
    IsAdmin  bool
}

func main() {
    generator := NewMockDataGenerator()
    user := generator.GenerateUserData()
    fmt.Printf("Generated User Data: %+v
", user)
}
