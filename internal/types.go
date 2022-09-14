package internal

type VisitorEvent struct {
	UserId            string `json:"user_id"`
	VisitorInternalId string `json:"visitor_internal_id"`
	HouseId           string `json:"house_id"`
	RegionId          string `json:"region_id"`
	DomainName        string `json:"domain_name"`
	EventType         string `json:"event_type"`
	EventCode         string `json:"event_code"`
	EventDate         string `json:"event_date"`
}
