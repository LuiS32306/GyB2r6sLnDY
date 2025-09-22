// 代码生成时间: 2025-09-22 15:24:34
package main

import (
    "context"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// WebContentGrabber is the main application struct
type WebContentGrabber struct {
    URLs []string
}

// NewWebContentGrabber creates a new instance of WebContentGrabber with the given URLs
func NewWebContentGrabber(urls []string) *WebContentGrabber {
    return &WebContentGrabber{URLs: urls}
}

// FetchContent fetches the content from each URL in the WebContentGrabber
func (w *WebContentGrabber) FetchContent() error {
    for _, url := range w.URLs {
        resp, err := http.Get(url)
        if err != nil {
            return fmt.Errorf("error fetching content from %s: %w", url, err)
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return fmt.Errorf("non-200 status code received from %s: %d", url, resp.StatusCode)
        }

        content, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return fmt.Errorf("error reading response body from %s: %w", url, err)
        }

        // Save the content to a file with the URL as the filename
        filename := fmt.Sprintf("%s.html", url)
        if err := ioutil.WriteFile(filename, content, 0644); err != nil {
            return fmt.Errorf("error writing content to file for %s: %w", url, err)
        }
    }
    return nil
}

// Run starts the web content grabber
func (w *WebContentGrabber) Run() {
    if err := w.FetchContent(); err != nil {
        fmt.Println("Error fetching content:", err)
        os.Exit(1)
    }
}

func main() {
    // Define the URLs to fetch content from
    urls := []string{
        "http://example.com",
        "http://golang.org",
    }

    // Create a new WebContentGrabber instance
    grabber := NewWebContentGrabber(urls)

    // Catch interrupts and gracefully exit on interrupt signal
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigs
        // Graceful shutdown
        fmt.Println("Shutting down...")
        os.Exit(0)
    }()

    // Run the grabber
    grabber.Run()
}
