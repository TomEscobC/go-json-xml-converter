package model

// Headers representa los metadatos del evento WHMG.
type Headers struct {
	Domain         string `json:"domain" xml:"Domain"`
	Channel        string `json:"channel" xml:"Channel"`
	Country        string `json:"country" xml:"Country"`
	EventId        string `json:"eventId" xml:"EventId"`
	Version        string `json:"version" xml:"Version"`
	Commerce       string `json:"commerce" xml:"Commerce"`
	Consumer       string `json:"consumer" xml:"Consumer"`
	Datetime       string `json:"datetime" xml:"Datetime"`
	EntityId       string `json:"entityId" xml:"EntityId"`
	MimeType       string `json:"mimeType" xml:"MimeType"`
	EventType      string `json:"eventType" xml:"EventType"`
	Timestamp      string `json:"timestamp" xml:"Timestamp"`
	Capability     string `json:"capability" xml:"Capability"`
	EntityType     string `json:"entityType" xml:"EntityType"`
	WhsEventSyncId string `json:"whsEventSyncId" xml:"WhsEventSyncId"`
}
