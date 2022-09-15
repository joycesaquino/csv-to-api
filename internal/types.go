package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type VisitorEvent struct {
	VisitorId          string    `json:"visitorId"`
	VisitorInternalId  string    `json:"externalId"`
	UserId             string    `json:"userId"`
	HouseId            string    `json:"houseId"`
	RegionId           string    `json:"regionId"`
	DomainName         string    `json:"domainName"`
	EventType          string    `json:"eventType"`
	EventCode          string    `json:"eventCode"`
	IdempotencyId      string    `json:"idempotencyId"`
	EventDate          time.Time `json:"eventDate"`
	VisitScheduledDate time.Time `json:"visitScheduledDate"`
}

func (v VisitorEvent) GetIdempotencyId() string {
	hash := md5.Sum([]byte(fmt.Sprint(v)))
	return hex.EncodeToString(hash[:])
}
