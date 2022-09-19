package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type VisitorEvent struct {
	VisitorId         string `json:"visitorId"`
	VisitorInternalId string `json:"externalId"`
	HouseId           string `json:"houseId"`
	RegionId          string `json:"regionId"`
	DomainName        string `json:"domainName"`
	EventType         string `json:"eventType"`
	EventCode         string `json:"eventCode"`
	IdempotencyId     string `json:"idempotencyId"`
	EventDate         string `json:"eventDate"`
}

type VisitorEventBody struct {
	VisitorId         string `json:"visitorId"`
	VisitorInternalId string `json:"externalId"`
	HouseId           int    `json:"houseId"`
	RegionId          int    `json:"regionId"`
	DomainName        string `json:"domainName"`
	EventType         string `json:"eventType"`
	EventCode         string `json:"eventCode"`
	IdempotencyId     string `json:"idempotencyId"`
	EventDate         string `json:"eventDate"`
}

func (v VisitorEvent) GetIdempotencyId() string {
	hash := md5.Sum([]byte(fmt.Sprint(v)))
	return hex.EncodeToString(hash[:])
}
