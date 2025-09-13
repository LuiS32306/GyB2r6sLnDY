// 代码生成时间: 2025-09-14 00:52:26
package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

// LogEntry 定义日志条目的结构
type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

// LogParserService 定义服务接口
type LogParserService struct {
	// UnimplementedLogParserService 可以添加其他方法
}

// ParseLogs 解析日志文件
func (s *LogParserService) ParseLogs(ctx context.Context, in *LogRequest) (*LogResponse, error) {
	if in == nil || in.FilePath == "" {
		return nil, fmt.Errorf("file path is required")
	}

	file, err := os.Open(in.FilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var entries []*LogEntry

	for scanner.Scan() {
		line := scanner.Text()
		entry, err := parseLine(line)
		if err != nil {
			log.Printf("error parsing line: %v", err)
			continue
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return &LogResponse{Entries: entries}, nil
}

// parseLine 解析单行日志
func parseLine(line string) (*LogEntry, error) {
	// 此处简化，实际解析逻辑可能更复杂
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid log line format")
	}

	// 假设日志格式为："时间 等级 消息"
	timestamp, err := time.Parse("2006-01-02 15:04:05", parts[0]+" "+parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to parse timestamp: %v", err)
	}

	return &LogEntry{
		Timestamp: timestamp,
		Level:     parts[2],
		Message:   strings.Join(parts[3:], " "),
	}, nil
}

// LogRequest 定义解析请求的结构
type LogRequest struct {
	FilePath string
}

// LogResponse 定义解析响应的结构
type LogResponse struct {
	Entries []*LogEntry
}

// The LogParserServer is the server API for LogParser service.
type LogParserServer struct {
	UnimplementedLogParserServer
}

func (s *LogParserServer) ParseLogs(ctx context.Context, req *LogRequest) (*LogResponse, error) {
	return NewLogParserService().ParseLogs(ctx, req)
}

// main 程序入口
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

ts := grpc.NewServer()
	RegisterLogParserServer(ts, &LogParserServer{})
	if err := ts.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
