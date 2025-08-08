package builder

import (
	"fmt"
	"strconv"

	"github.com/TomEscobC/go-json-xml-converter/internal/model"
)

// DistributionOrderBuilder implementa el patrón Builder para construir objetos DistributionOrder.
type DistributionOrderBuilder struct {
	order *model.DistributionOrder
}

// NewDistributionOrderBuilder crea una nueva instancia de DistributionOrderBuilder.
func NewDistributionOrderBuilder() *DistributionOrderBuilder {
	return &DistributionOrderBuilder{
		order: &model.DistributionOrder{},
	}
}

// SetBasicInfo establece la información básica del pedido.
func (b *DistributionOrderBuilder) SetBasicInfo(data map[string]interface{}) *DistributionOrderBuilder {
	if data == nil {
		return b
	}

	if distributionOrderId, ok := data["distributionOrderId"].(string); ok {
		b.order.DistributionOrderId = distributionOrderId
	}
	if orderType, ok := data["orderType"].(string); ok {
		b.order.OrderType = orderType
	}
	if orderedDateTime, ok := data["orderedDateTime"].(string); ok {
		b.order.OrderedDttm = orderedDateTime
	}
	if doFulfillmentStatus, ok := data["doFulfillmentStatus"].(string); ok {
		b.order.DoFulfillmentStatus = doFulfillmentStatus
	}
	if isCancelled, ok := data["isCancelled"].(string); ok {
		b.order.IsCancelled = isCancelled
	}
	// businessUnit puede ser número o string en el JSON
	if businessUnit, ok := data["businessUnit"]; ok {
		b.order.BusinessUnit = fmt.Sprintf("%v", businessUnit)
	}
	if referenceField3, ok := data["referenceField3"].(string); ok {
		b.order.ReferenceField3 = referenceField3
	}
	if referenceField4, ok := data["referenceField4"].(string); ok {
		b.order.ReferenceField4 = referenceField4
	}
	if routeTo, ok := data["routeTo"].(string); ok {
		b.order.RouteTo = routeTo
	}
	if orderType, ok := data["type"].(string); ok {
		b.order.DistributionOrderType = orderType
	}

	return b
}

// SetTimeInfo establece la información de tiempos.
func (b *DistributionOrderBuilder) SetTimeInfo(data map[string]interface{}) *DistributionOrderBuilder {
	if data == nil {
		return b
	}

	if pickupStartDateTime, ok := data["pickupStartDateTime"].(string); ok {
		b.order.PickupStartDttm = pickupStartDateTime
	}
	if pickupEndDateTime, ok := data["pickupEndDateTime"].(string); ok {
		b.order.PickupEndDttm = pickupEndDateTime
	}
	if deliveryStartDateTime, ok := data["deliveryStartDateTime"].(string); ok {
		b.order.DeliveryStartDttm = deliveryStartDateTime
	}
	if deliveryEndDateTime, ok := data["deliveryEndDateTime"].(string); ok {
		b.order.DeliveryEndDttm = deliveryEndDateTime
	}

	return b
}

// SetLabelInfo establece la información de etiquetas.
func (b *DistributionOrderBuilder) SetLabelInfo(data map[string]interface{}) *DistributionOrderBuilder {
	if data == nil {
		return b
	}

	// lpnCubingIndicator puede ser número o string
	if lpnCubingIndicator, ok := data["lpnCubingIndicator"]; ok {
		b.order.LpnCubingIndicator = fmt.Sprintf("%v", lpnCubingIndicator)
	}
	if partialShipConfFlag, ok := data["partialShipConfFlag"].(string); ok {
		b.order.PartlShipConfFlag = partialShipConfFlag
	}
	if contentLabelType, ok := data["contentLabelType"].(string); ok {
		b.order.ContentLabelType = contentLabelType
	}
	if lpnLabelType, ok := data["lpnLabelType"].(string); ok {
		b.order.LpnLabelType = lpnLabelType
	}

	return b
}

// SetFacilityInfo establece la información de facilidades.
func (b *DistributionOrderBuilder) SetFacilityInfo(data map[string]interface{}) *DistributionOrderBuilder {
	if data == nil {
		return b
	}

	if originFacilityAliasId, ok := data["originFacilityAliasId"].(string); ok {
		b.order.OriginFacilityAliasId = originFacilityAliasId
	}
	if destinationFacilityAliasId, ok := data["destinationFacilityAliasId"].(string); ok {
		b.order.DestinationFacilityAliasId = destinationFacilityAliasId
	}

	return b
}

// SetLineItems establece los items de línea.
func (b *DistributionOrderBuilder) SetLineItems(items []interface{}) *DistributionOrderBuilder {
	if items == nil {
		b.order.LineItems = []model.LineItem{}
		return b
	}

	var lineItems []model.LineItem
	for _, item := range items {
		if itemMap, ok := item.(map[string]interface{}); ok {
			lineItem := b.buildLineItem(itemMap)
			lineItems = append(lineItems, lineItem)
		}
	}
	b.order.LineItems = lineItems

	return b
}

// Build construye y retorna el objeto DistributionOrder.
func (b *DistributionOrderBuilder) Build() *model.DistributionOrder {
	// Asegurar que LineItems no sea nil
	if b.order.LineItems == nil {
		b.order.LineItems = []model.LineItem{}
	}
	return b.order
}

// buildLineItem construye un LineItem individual.
func (b *DistributionOrderBuilder) buildLineItem(data map[string]interface{}) model.LineItem {
	return model.LineItem{
		Name:                   getStringFromMap(data, "name"),
		DoLineNumber:           getStringFromMap(data, "doLineNumber"),
		FulfillmentType:        getStringFromMap(data, "fulfillmentType"),
		StoreDepartment:        getStringFromMap(data, "storeDepartment"),
		MerchandizeGroup:       getStringFromMap(data, "merchandizeGroup"),
		AllocationSourceId:     getStringFromMap(data, "allocationSourceId"),
		AllocationSourceType:   getStringFromMap(data, "allocationSourceType"),
		AllocationSourceLineId: getStringFromMap(data, "allocationSourceLineId"),
		ShelfDays:              getIntFromMap(data, "shelfDays"),
		Weight:                 b.buildWeight(getMapFromMap(data, "weight")),
		Quantity:               b.buildQuantity(getMapFromMap(data, "quantity")),
		InventoryAttributes:    b.buildInventoryAttributes(getMapFromMap(data, "inventoryAttributes")),
	}
}

// buildWeight construye un objeto Weight.
func (b *DistributionOrderBuilder) buildWeight(data map[string]interface{}) model.Weight {
	if data == nil {
		return model.Weight{}
	}
	return model.Weight{
		Uom:     getStringFromMap(data, "uom"),
		Planned: getFloat64FromMap(data, "planned"),
	}
}

// buildQuantity construye un objeto Quantity.
func (b *DistributionOrderBuilder) buildQuantity(data map[string]interface{}) model.Quantity {
	if data == nil {
		return model.Quantity{}
	}
	return model.Quantity{
		Uom:   getStringFromMap(data, "uom"),
		Order: getIntFromMap(data, "order"),
	}
}

// buildInventoryAttributes construye un objeto InventoryAttributes.
func (b *DistributionOrderBuilder) buildInventoryAttributes(data map[string]interface{}) model.InventoryAttributes {
	if data == nil {
		return model.InventoryAttributes{}
	}
	return model.InventoryAttributes{
		BatchNumber:     getStringFromMap(data, "batchNumber"),
		InventoryType:   getStringFromMap(data, "inventoryType"),
		CountryOfOrigin: getStringFromMap(data, "countryOfOrigin"),
	}
}

// Funciones auxiliares para obtener valores de mapas de forma segura.

func getStringFromMap(data map[string]interface{}, key string) string {
	if val, exists := data[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
		// Si no es string, convertir a string
		return fmt.Sprintf("%v", val)
	}
	return ""
}

func getIntFromMap(data map[string]interface{}, key string) int {
	if val, exists := data[key]; exists {
		switch v := val.(type) {
		case float64: // JSON numbers are float64 by default
			return int(v)
		case int:
			return v
		case int64:
			return int(v)
		case string:
			// Intentamos convertir string a int
			if i, err := strconv.Atoi(v); err == nil {
				return i
			}
		}
	}
	return 0
}

func getFloat64FromMap(data map[string]interface{}, key string) float64 {
	if val, exists := data[key]; exists {
		switch v := val.(type) {
		case float64:
			return v
		case int:
			return float64(v)
		case int64:
			return float64(v)
		case string:
			// Intentamos convertir string a float64
			if f, err := strconv.ParseFloat(v, 64); err == nil {
				return f
			}
		}
	}
	return 0.0
}

func getMapFromMap(data map[string]interface{}, key string) map[string]interface{} {
	if val, exists := data[key]; exists {
		if mapVal, ok := val.(map[string]interface{}); ok {
			return mapVal
		}
	}
	return nil
}

// GetArrayFromMap es una función auxiliar pública para obtener arrays del mapa de datos.
func GetArrayFromMap(data map[string]interface{}, key string) []interface{} {
	if val, exists := data[key]; exists {
		if arr, ok := val.([]interface{}); ok {
			return arr
		}
	}
	return []interface{}{}
}
