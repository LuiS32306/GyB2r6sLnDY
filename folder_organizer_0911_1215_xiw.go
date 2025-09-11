// 代码生成时间: 2025-09-11 12:15:30
package main

import (
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

// FolderOrganizer is the service that handles folder organization logic
type FolderOrganizer struct{}

// OrganizeDirectory takes a directory path and organizes it
// by moving files into folders named after their file extensions.
func (f *FolderOrganizer) OrganizeDirectory(path string) error {
    files, err := ioutil.ReadDir(path)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        ext := strings.TrimPrefix(filepath.Ext(file.Name()), ".")
        if ext == "" {
            continue
        }

        // Create a new directory for the file extension if it doesn't exist
        dirPath := filepath.Join(path, ext)
        if _, err := os.Stat(dirPath); os.IsNotExist(err) {
            if err := os.MkdirAll(dirPath, 0755); err != nil {
                return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
            }
        }

        // Move the file into the new directory
        srcPath := filepath.Join(path, file.Name())
        destPath := filepath.Join(dirPath, file.Name())
        if err := os.Rename(srcPath, destPath); err != nil {
            return fmt.Errorf("failed to move file %s to %s: %w", srcPath, destPath, err)
        }
    }

    return nil
}

func main() {
    // Example usage of FolderOrganizer
    path := "./example_directory"
    organizer := &FolderOrganizer{}

    if err := organizer.OrganizeDirectory(path); err != nil {
        log.Fatalf("Error organizing directory: %s", err)
    }
    fmt.Println("Directory organized successfully.")
}
