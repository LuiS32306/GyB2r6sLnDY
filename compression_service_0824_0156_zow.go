// 代码生成时间: 2025-08-24 01:56:17
package main

import (
    "archive/zip"
    "bufio"
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net"
    "os"
    "path/filepath"
)

// CompressionService defines the gRPC service for compression and decompression.
type CompressionService struct {
    // no fields needed for this example
}

// CompressFile compresses a single file into a zip archive.
func (s *CompressionService) CompressFile(stream CompressionService_CompressFileServer) error {
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        // Process the received file and add it to the zip archive.
        if err := addFilesToZip(req.GetFilePath(), req.GetData(), stream); err != nil {
            return err
        }
    }
    return nil
}

// DecompressFile decompresses a zip archive and returns the contents.
func (s *CompressionService) DecompressFile(stream CompressionService_DecompressFileServer) error {
    zipData, err := ioutil.ReadAll(stream)
    if err != nil {
        return err
    }

    // Create a buffer to store the decompressed data.
    buffer := new(bytes.Buffer)

    // Create a new zip reader.
    reader, err := zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
    if err != nil {
        return err
    }

    // Iterate through the files in the zip archive.
    for _, file := range reader.File {
        rc, err := file.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        // Copy the file data to the buffer.
        if _, err := io.Copy(buffer, rc); err != nil {
            return err
        }
    }

    // Send the decompressed data back to the client.
    if err := stream.SendAndClose(&DecompressedFile{Data: buffer.Bytes()}); err != nil {
        return err
        }
    return nil
}

// addFilesToZip adds files to a zip archive.
func addFilesToZip(filePath string, data []byte, stream CompressionService_CompressFileServer) error {
    // Create a buffer to store the zip archive.
    buffer := new(bytes.Buffer)
    writer := zip.NewWriter(buffer)

    // Create a file within the zip archive.
    f, err := writer.Create(filePath)
    if err != nil {
        return err
    }

    // Write the file data to the archive.
    if _, err := f.Write(data); err != nil {
        return err
    }

    // Close the zip writer to finish the archive.
    if err := writer.Close(); err != nil {
        return err
    }

    // Send the zip archive back to the client.
    if err := stream.Send(&CompressedFile{Data: buffer.Bytes()}); err != nil {
        return err
    }

    return nil
}

// main function to start the gRPC server.
func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    fmt.Println("Listening on port 50051")

    s := grpc.NewServer()
    RegisterCompressionServiceServer(s, &CompressionService{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
