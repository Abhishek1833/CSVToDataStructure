package main

import (
    "encoding/csv"
    "fmt"
    "os"
)

type Record struct {
    Name    string
    Age     int
    Address string
}

func main() {
    // Open the CSV file
    f, err := os.Open("data.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Create a new CSV reader
    reader := csv.NewReader(f)

    // Read the CSV file line by line
    lines, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Parse the CSV data into a slice of structs
    var records []Record
    for _, line := range lines {
        age := 0
        fmt.Sscanf(line[1], "%d", &age) // convert string to int
        records = append(records, Record{
            Name:    line[0],
            Age:     age,
            Address: line[2],
        })
    }

    // Output the data as a table
    fmt.Println("Name\tAge\tAddress")
    for _, record := range records {
        fmt.Printf("%s\t%d\t%s\n", record.Name, record.Age, record.Address)
    }
}
