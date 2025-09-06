// 代码生成时间: 2025-09-07 07:10:23
package main

import (
    "bytes"
    "encoding/csv"
    "io"
    "log"
    "os"
    "strconv"
)

// CSVRow represents a row in the CSV file
type CSVRow struct {
    Field1 string
    Field2 string
    Field3 string
    // Add more fields as needed
}

// ProcessCSVFile processes a single CSV file
func ProcessCSVFile(filename string) ([]CSVRow, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    rows, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var csvRows []CSVRow
    for _, row := range rows {
        if len(row) < 3 {
            continue // Skip rows with fewer than 3 columns
        }
        // Convert row to CSVRow structure, assuming first three fields are needed
        csvRows = append(csvRows, CSVRow{
            Field1: row[0],
            Field2: row[1],
            Field3: row[2],
        })
    }
    return csvRows, nil
}

// ProcessCSVFiles processes multiple CSV files
func ProcessCSVFiles(files []string) ([]CSVRow, error) {
    var allRows []CSVRow
    for _, filename := range files {
        csvRows, err := ProcessCSVFile(filename)
        if err != nil {
            log.Printf("Error processing file %s: %v", filename, err)
            continue
        }
        allRows = append(allRows, csvRows...)
    }
    return allRows, nil
}

func main() {
    // List of CSV files to process
    filenames := []string{
        "file1.csv",
        "file2.csv",
        // Add more file names as needed
    }
    
    csvRows, err := ProcessCSVFiles(filenames)
    if err != nil {
        log.Fatal(err)
    }
    
    // Do something with the processed CSV rows
    // For example, print them to the console
    for _, row := range csvRows {
        log.Printf("Processed row: Field1=%s, Field2=%s, Field3=%s", row.Field1, row.Field2, row.Field3)
    }
}
