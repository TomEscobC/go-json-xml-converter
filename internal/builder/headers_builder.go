package builder

import (
	"github.com/TomEscobC/go-json-xml-converter/internal/model"
)

// HeadersBuilder implementa el patrón Builder para construir objetos Headers.
type HeadersBuilder struct {
	headers *model.Headers
}

// NewHeadersBuilder crea una nueva instancia de HeadersBuilder.
func NewHeadersBuilder() *HeadersBuilder {
	return &HeadersBuilder{
		headers: &model.Headers{},
	}
}

// SetDomain establece el dominio.
func (b *HeadersBuilder) SetDomain(domain string) *HeadersBuilder {
	b.headers.Domain = domain
	return b
}

// SetChannel establece el canal.
func (b *HeadersBuilder) SetChannel(channel string) *HeadersBuilder {
	b.headers.Channel = channel
	return b
}

// SetCountry establece el país.
func (b *HeadersBuilder) SetCountry(country string) *HeadersBuilder {
	b.headers.Country = country
	return b
}

// SetEventId establece el ID del evento.
func (b *HeadersBuilder) SetEventId(eventId string) *HeadersBuilder {
	b.headers.EventId = eventId
	return b
}

// SetVersion establece la versión.
func (b *HeadersBuilder) SetVersion(version string) *HeadersBuilder {
	b.headers.Version = version
	return b
}

// SetCommerce establece el comercio.
func (b *HeadersBuilder) SetCommerce(commerce string) *HeadersBuilder {
	b.headers.Commerce = commerce
	return b
}

// SetConsumer establece el consumidor.
func (b *HeadersBuilder) SetConsumer(consumer string) *HeadersBuilder {
	b.headers.Consumer = consumer
	return b
}

// SetDatetime establece la fecha y hora.
func (b *HeadersBuilder) SetDatetime(datetime string) *HeadersBuilder {
	b.headers.Datetime = datetime
	return b
}

// SetEntityId establece el ID de entidad.
func (b *HeadersBuilder) SetEntityId(entityId string) *HeadersBuilder {
	b.headers.EntityId = entityId
	return b
}

// SetMimeType establece el tipo MIME.
func (b *HeadersBuilder) SetMimeType(mimeType string) *HeadersBuilder {
	b.headers.MimeType = mimeType
	return b
}

// SetEventType establece el tipo de evento.
func (b *HeadersBuilder) SetEventType(eventType string) *HeadersBuilder {
	b.headers.EventType = eventType
	return b
}

// SetTimestamp establece el timestamp.
func (b *HeadersBuilder) SetTimestamp(timestamp string) *HeadersBuilder {
	b.headers.Timestamp = timestamp
	return b
}

// SetCapability establece la capacidad.
func (b *HeadersBuilder) SetCapability(capability string) *HeadersBuilder {
	b.headers.Capability = capability
	return b
}

// SetEntityType establece el tipo de entidad.
func (b *HeadersBuilder) SetEntityType(entityType string) *HeadersBuilder {
	b.headers.EntityType = entityType
	return b
}

// SetWhsEventSyncId establece el ID de sincronización del evento WHS.
func (b *HeadersBuilder) SetWhsEventSyncId(whsEventSyncId string) *HeadersBuilder {
	b.headers.WhsEventSyncId = whsEventSyncId
	return b
}

// Build construye y retorna el objeto Headers.
func (b *HeadersBuilder) Build() *model.Headers {
	return b.headers
}

// SetAllHeaders establece todos los headers desde un mapa.
// Esto nos facilita la construcción desde el JSON de entrada.
func (b *HeadersBuilder) SetAllHeaders(headersMap map[string]interface{}) *HeadersBuilder {
	if headersMap == nil {
		return b
	}

	if domain, ok := headersMap["domain"].(string); ok {
		b.SetDomain(domain)
	}
	if channel, ok := headersMap["channel"].(string); ok {
		b.SetChannel(channel)
	}
	if country, ok := headersMap["country"].(string); ok {
		b.SetCountry(country)
	}
	if eventId, ok := headersMap["eventId"].(string); ok {
		b.SetEventId(eventId)
	}
	if version, ok := headersMap["version"].(string); ok {
		b.SetVersion(version)
	}
	if commerce, ok := headersMap["commerce"].(string); ok {
		b.SetCommerce(commerce)
	}
	if consumer, ok := headersMap["consumer"].(string); ok {
		b.SetConsumer(consumer)
	}
	if datetime, ok := headersMap["datetime"].(string); ok {
		b.SetDatetime(datetime)
	}
	if entityId, ok := headersMap["entityId"].(string); ok {
		b.SetEntityId(entityId)
	}
	if mimeType, ok := headersMap["mimeType"].(string); ok {
		b.SetMimeType(mimeType)
	}
	if eventType, ok := headersMap["eventType"].(string); ok {
		b.SetEventType(eventType)
	}
	if timestamp, ok := headersMap["timestamp"].(string); ok {
		b.SetTimestamp(timestamp)
	}
	if capability, ok := headersMap["capability"].(string); ok {
		b.SetCapability(capability)
	}
	if entityType, ok := headersMap["entityType"].(string); ok {
		b.SetEntityType(entityType)
	}
	if whsEventSyncId, ok := headersMap["whsEventSyncId"].(string); ok {
		b.SetWhsEventSyncId(whsEventSyncId)
	}

	return b
}
