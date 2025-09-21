// 代码生成时间: 2025-09-22 03:32:52
package main

import (
    "archive/zip"
    "compress/gzip"
    "flag"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"
)

// 工具主要功能：压缩文件解压工具

// Decompressor 定义一个解压缩接口
type Decompressor interface {
    Decompress(src, dest string) error
}

// ZipDecompressor 结构体实现 Zip 解压缩
type ZipDecompressor struct{}

// Decompress 实现 Zip 解压缩功能
func (d ZipDecompressor) Decompress(src, dest string) error {
    reader, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer reader.Close()

    for _, file := range reader.File {
        filePath := filepath.Join(dest, file.Name)
        if file.FileInfo().IsDir() {
            // 创建文件夹
            os.MkdirAll(filePath, os.ModePerm)
            continue
        }

        // 打开文件并复制内容
        fileReader, err := file.Open()
        if err != nil {
            return err
        }
        defer fileReader.Close()

        targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.FileInfo().Mode())
        if err != nil {
            return err
        }
        defer targetFile.Close()

        _, err = io.Copy(targetFile, fileReader)
        if err != nil {
            return err
        }
    }
    return nil
}

// GzipDecompressor 结构体实现 Gzip 解压缩
type GzipDecompressor struct{}

// Decompress 实现 Gzip 解压缩功能
func (d GzipDecompressor) Decompress(src, dest string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    destFile, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer destFile.Close()

    gzReader, err := gzip.NewReader(srcFile)
    if err != nil {
        return err
    }
    defer gzReader.Close()

    _, err = io.Copy(destFile, gzReader)
    if err != nil {
        return err
    }
    return nil
}

func main() {
    // 解析命令行参数
    var compressType string
    var src, dest string
    flag.StringVar(&compressType, "type", "", "Compression type (zip or gzip)")
    flag.StringVar(&src, "src", "", "Source file path")
    flag.StringVar(&dest, "dest", "", "Destination directory path")
    flag.Parse()

    if src == "" || dest == "" || compressType == "" {
        fmt.Println("Usage: compress_utility -type [zip|gzip] -src <source file> -dest <destination directory>")
        return
    }

    // 根据压缩类型选择不同的解压缩器
    var decompressor Decompressor
    switch strings.ToLower(compressType) {
    case "zip":
        decompressor = ZipDecompressor{}
    case "gzip":
        decompressor = GzipDecompressor{}
    default:
        fmt.Println("Unsupported compression type")
        return
    }

    // 解压缩文件
    if err := decompressor.Decompress(src, dest); err != nil {
        fmt.Printf("Decompression failed: %s
", err)
        return
    }
    fmt.Println("Decompression completed successfully")
}