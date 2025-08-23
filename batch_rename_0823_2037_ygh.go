// 代码生成时间: 2025-08-23 20:37:47
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

// RenameFile represents a file and its new name
type RenameFile struct {
    Source string `json:"source"`
    Dest   string `json:"dest"`
}

// BatchRenameService provides methods for batch renaming files
type BatchRenameService struct{}

// Rename renames a list of files
func (s *BatchRenameService) Rename(files []RenameFile) error {
    for _, file := range files {
        if err := renameFile(file.Source, file.Dest); err != nil {
            return err
        }
    }
    return nil
}

// renameFile performs the actual file renaming operation
func renameFile(src, dest string) error {
    // Check if the source file exists
    if _, err := os.Stat(src); os.IsNotExist(err) {
        return fmt.Errorf("source file does not exist: %s", src)
    }

    // Ensure the destination path is not already taken
    if _, err := os.Stat(dest); !os.IsNotExist(err) {
        return fmt.Errorf("destination file already exists: %s", dest)
    }

    // Perform the rename operation
    if err := os.Rename(src, dest); err != nil {
        return fmt.Errorf("failed to rename %s to %s: %w", src, dest, err)
    }

    return nil
}

func main() {
    // Example usage of BatchRenameService
    service := BatchRenameService{}

    // List of files to rename
    files := []RenameFile{
        {Source: "file1.txt", Dest: "newfile1.txt"},
        {Source: "file2.txt", Dest: "newfile2.txt"},
    }

    // Attempt to rename the files
    if err := service.Rename(files); err != nil {
        log.Fatalf("Error renaming files: %v", err)
    }
    fmt.Println("Files renamed successfully.")
}
