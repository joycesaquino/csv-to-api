package internal

import (
	"encoding/json"
	"log"
	"reflect"
)

type Converter struct {
}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) CsvToVisitorEvents(messages [][]string) ([]VisitorEvent, error) {
	var visitorEvent VisitorEvent
	var outputs []VisitorEvent

	visitorEventObject := make(map[string]interface{}, reflect.ValueOf(visitorEvent).NumField())

	headersArr := []string{"visitorId", "externalId", "houseId", "regionId", "domainName", "eventType", "eventCode", "eventDate"}

	log.Printf("Starting converter messages to output.")

	for _, message := range messages {
		for i, value := range message {
			visitorEventObject[headersArr[i]] = value
		}

		object, err := json.Marshal(visitorEventObject)
		if err != nil {
			log.Printf("Error on marshal interface %s :", err)
		}

		err = json.Unmarshal(object, &visitorEvent)
		if err != nil {
			log.Printf("Error on unmarshal interface to visitor event json %s :", err)
		}

		visitorEvent.IdempotencyId = visitorEvent.GetIdempotencyId()
		outputs = append(outputs, visitorEvent)
	}

	log.Printf("End converter messages to output.")
	return outputs, nil
}
