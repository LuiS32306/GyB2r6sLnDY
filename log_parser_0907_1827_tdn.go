// 代码生成时间: 2025-09-07 18:27:44
// log_parser.go

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// LogEntry represents a single log entry with its timestamp and message.
type LogEntry struct {
	Timestamp string
	Message   string
}

// ParseLog parses a log file and returns a slice of LogEntry objects.
func ParseLog(filePath string) ([]LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var entries []LogEntry
	scanner := newLineScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		entry, err := parseLine(line)
		if err != nil {
			log.Printf("ignoring invalid log entry: %s", line)
			continue
		}
		entries = append(entries, entry)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return entries, nil
}

// parseLine takes a log line and parses it into a LogEntry.
func parseLine(line string) (LogEntry, error) {
	// Assuming the log format is `timestamp message`
	parts := strings.SplitN(line, " ", 2)
	if len(parts) != 2 {
		return LogEntry{}, fmt.Errorf("invalid log line format: %s", line)
	}
	return LogEntry{Timestamp: parts[0], Message: parts[1]}, nil
}

// newLineScanner creates a Scanner that reads lines from an io.Reader.
func newLineScanner(reader io.Reader) *strings.Scanner {
	return strings.NewScanner(reader)
}

func main() {
	// Example usage of ParseLog function.
	entries, err := ParseLog("example.log")
	if err != nil {
		log.Fatalf("error parsing log file: %s", err)
	}
	for _, entry := range entries {
		fmt.Printf("Timestamp: %s, Message: %s
", entry.Timestamp, entry.Message)
	}
}