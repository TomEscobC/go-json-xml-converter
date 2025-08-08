package model

// DistributionOrder representa la estructura del cuerpo del evento WHMG.
// Corresponde a la estructura del body en el JSON de entrada.
type DistributionOrder struct {
	XMLName                    string     `xml:"DistributionOrder"`
	DistributionOrderId        string     `xml:"DistributionOrderId"`
	OrderType                  string     `xml:"OrderType"`
	OrderedDttm                string     `xml:"OrderedDttm"`
	DoFulfillmentStatus        string     `xml:"DoFulfillmentStatus"`
	IsCancelled                string     `xml:"IsCancelled"`
	BusinessUnit               string     `xml:"BusinessUnit"`
	ReferenceField3            string     `xml:"ReferenceField3"`
	ReferenceField4            string     `xml:"ReferenceField4"`
	PickupStartDttm            string     `xml:"PickupStartDttm"`
	PickupEndDttm              string     `xml:"PickupEndDttm"`
	DeliveryStartDttm          string     `xml:"DeliveryStartDttm"`
	DeliveryEndDttm            string     `xml:"DeliveryEndDttm"`
	LpnCubingIndicator         string     `xml:"LpnCubingIndicator"`
	PartlShipConfFlag          string     `xml:"PartlShipConfFlag"`
	ContentLabelType           string     `xml:"ContentLabelType"`
	LpnLabelType               string     `xml:"LpnLabelType"`
	OriginFacilityAliasId      string     `xml:"OriginFacilityAliasId"`
	DestinationFacilityAliasId string     `xml:"DestinationFacilityAliasId"`
	RouteTo                    string     `xml:"RouteTo"`
	DistributionOrderType      string     `xml:"DistributionOrderType"`
	LineItems                  []LineItem `xml:"LineItems>LineItem"`
}

// LineItem representa un ítem de línea en el pedido de distribución.
type LineItem struct {
	Name                   string              `xml:"Name"`
	DoLineNumber           string              `xml:"DoLineNumber"`
	FulfillmentType        string              `xml:"FulfillmentType"`
	StoreDepartment        string              `xml:"StoreDepartment"`
	MerchandizeGroup       string              `xml:"MerchandizeGroup"`
	AllocationSourceId     string              `xml:"AllocationSourceId"`
	AllocationSourceType   string              `xml:"AllocationSourceType"`
	AllocationSourceLineId string              `xml:"AllocationSourceLineId"`
	ShelfDays              int                 `xml:"ShelfDays"`
	Weight                 Weight              `xml:"Weight"`
	Quantity               Quantity            `xml:"Quantity"`
	InventoryAttributes    InventoryAttributes `xml:"InventoryAttributes"`
}

// Weight representa el peso de un ítem.
type Weight struct {
	Uom     string  `xml:"Uom"`
	Planned float64 `xml:"Planned"`
}

// Quantity representa la cantidad de un ítem.
type Quantity struct {
	Uom   string `xml:"Uom"`
	Order int    `xml:"Order"`
}

// InventoryAttributes representa los atributos de inventario de un ítem.
type InventoryAttributes struct {
	BatchNumber     string `xml:"BatchNumber"`
	InventoryType   string `xml:"InventoryType"`
	CountryOfOrigin string `xml:"CountryOfOrigin"`
}
