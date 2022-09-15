package main

import (
	"context"
	"csv-to-api/internal"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	csvFile, err := os.Open("internal/events-test.csv")
	var messages [][]string
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Print(err)
	}
	messages = append(messages, records...)

	converter := internal.NewConverter()
	httpClient := internal.NewClient()

	visitorEvents, err := converter.CsvToVisitorEvents(messages)
	if err != nil {
		return
	}

	for _, event := range visitorEvents {
		err := httpClient.Post(context.Background(), event)
		if err != nil {
			return
		}
	}

}
