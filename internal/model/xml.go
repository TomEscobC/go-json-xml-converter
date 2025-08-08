package model

import "encoding/xml"

type XMLHeader struct {
	Domain         string `json:"domain"`
	Channel        string `json:"channel"`
	Country        string `json:"country"`
	EventID        string `json:"eventId"`
	Version        string `json:"version"`
	Commerce       string `json:"commerce"`
	Consumer       string `json:"consumer"`
	Datetime       string `json:"datetime"`
	EntityID       string `json:"entityId"`
	MimeType       string `json:"mimeType"`
	EventType      string `json:"eventType"`
	Timestamp      string `json:"timestamp"`
	Capability     string `json:"capability"`
	EntityType     string `json:"entityType"`
	WhsEventSyncID string `json:"whsEventSyncId"`
}

type DistributionOrderXML struct {
	XMLName xml.Name `xml:"tXml"`
	Header  XMLHeader

	Message struct {
		Order struct {
			DistributionOrderId   string `xml:"DistributionOrderId"`
			OrderType             string `xml:"OrderType"`
			OrderedDttm           string `xml:"OrderedDttm"`
			DoFulfillmentStatus   string `xml:"DoFulfillmentStatus"`
			IsCancelled           string `xml:"IsCancelled"`
			Priority              int    `xml:"Priority"`
			ReferenceField3       string `xml:"ReferenceField3"`
			ReferenceField4       string `xml:"ReferenceField4"`
			PickupStartDttm       string `xml:"PickupStartDttm"`
			PickupEndDttm         string `xml:"PickupEndDttm"`
			DeliveryStartDttm     string `xml:"DeliveryStartDttm"`
			DeliveryEndDttm       string `xml:"DeliveryEndDttm"`
			RouteTo               string `xml:"RouteTo"`
			OriginFacilityId      string `xml:"OriginFacilityAliasId"`
			DestinationFacility   string `xml:"DestinationFacilityAliasId"`
			LpnLabelType          string `xml:"LpnLabelType"`
			ContentLabelType      string `xml:"ContentLabelType"`
			PartlShipConfFlag     string `xml:"PartlShipConfFlag"`
			DistributionOrderType string `xml:"DistributionOrderType"`
		} `xml:"DistributionOrder"`
	} `xml:"Message"`
}
