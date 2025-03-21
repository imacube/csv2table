package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	// Check for the CSV file path argument.
	if len(os.Args) < 2 {
		log.Fatal("Usage: csv2table <csv file>")
	}

	// Open the CSV file.
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	// Read the CSV file.
	csvFile := csv.NewReader(file)
	data, err := csvFile.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	if len(data) == 0 {
		log.Fatal("CSV file is empty")
	}

	table := tablewriter.NewWriter(os.Stdout)

	// Use the first row of the CSV as the header.
	table.SetHeader(data[0])

	// Keep all columns to a single line
	table.SetAutoWrapText(false)

	// Add a line delineating each row
	table.SetRowLine(true)

	// Append the rest of the CSV rows to the table.
	for _, row := range data[1:] {
		table.Append(row)
	}

	// Render the table to the console.
	table.Render()
}
