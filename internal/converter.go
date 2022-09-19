package internal

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
)

type Converter struct {
}

func NewConverter() *Converter {
	return &Converter{}
}

func (c *Converter) CsvToVisitorEvents(messages [][]string) ([]VisitorEventBody, error) {
	var visitorEvent VisitorEvent
	var outputs []VisitorEventBody

	visitorEventObject := make(map[string]interface{}, reflect.ValueOf(visitorEvent).NumField())

	headersArr := []string{"externalId", "visitorId", "houseId", "regionId", "domainName", "eventType", "eventCode", "eventDate"}

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
		houseId, err := strconv.ParseInt(visitorEvent.HouseId, 10, 64)
		regionId, err := strconv.ParseInt(visitorEvent.RegionId, 10, 64)

		if err != nil {
			log.Printf("Error converting visitor event to body %s", err)
		}

		outputs = append(outputs, VisitorEventBody{
			VisitorId:         visitorEvent.VisitorId,
			VisitorInternalId: visitorEvent.VisitorInternalId,
			HouseId:           int(houseId),
			RegionId:          int(regionId),
			DomainName:        visitorEvent.DomainName,
			EventType:         visitorEvent.EventType,
			EventCode:         visitorEvent.EventCode,
			IdempotencyId:     visitorEvent.IdempotencyId,
			EventDate:         visitorEvent.EventDate,
		})
	}

	log.Printf("End converter messages to output.")
	return outputs, nil
}
