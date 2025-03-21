package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	// Parse arguments
	setRowLine := flag.Bool("rowline", true, "Draws a line between each row")
	header := flag.Bool("header", true, "Use the first row as the header")
	flag.Parse()

	// Check for the CSV file path argument
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: csv2table [arguments] [csv file]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Open the CSV file
	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	// Read the CSV file
	csvFile := csv.NewReader(file)
	data, err := csvFile.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}
	if len(data) == 0 {
		log.Fatal("CSV file is empty")
	}

	table := tablewriter.NewWriter(os.Stdout)

	// Should the first row be used as the table header?
	firstRow := 0
	if *header {
		table.SetHeader(data[0])
		firstRow = 1
	}

	// Keep all columns to a single line
	table.SetAutoWrapText(false)

	// Add a line delineating each row
	table.SetRowLine(*setRowLine)

	// Append the rest of the CSV rows to the table
	for _, row := range data[firstRow:] {
		table.Append(row)
	}

	// Render the table to the console
	table.Render()
}
