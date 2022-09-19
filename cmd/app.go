package main

import (
	"context"
	"csv-to-api/internal"
	"encoding/csv"
	"fmt"
	"log"
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
	//user_id,user_internal_id,house_id,region_id,domain_name,event_type,event_code,event_date

	for _, event := range visitorEvents {
		err := httpClient.Post(context.Background(), event)
		if err != nil {
			log.Printf("Error on send visitor event %v to Hubs API - Error: %s\n", event, err)
		}
	}

}
