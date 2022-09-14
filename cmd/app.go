package main

import (
	"csv-to-api/internal"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	csvFile, err := os.Open("visitor-event-test.csv")
	var messages [][]string
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Print(err)
	}
	messages = append(messages, records...)

	converter := internal.NewConverter()
	visitorEvents, err := converter.CsvToPriceHistory(messages)
	if err != nil {
		return
	}
	log.Println(visitorEvents)

}
