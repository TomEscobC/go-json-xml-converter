package facade

import (
	"fmt"

	"github.com/TomEscobC/go-json-xml-converter/internal/builder"
	"github.com/TomEscobC/go-json-xml-converter/internal/model"
)

// ConverterFacade implementa el patrón Facade para simplificar la conversión
type ConverterFacade struct {
	headersBuilder *builder.HeadersBuilder
	orderBuilder   *builder.DistributionOrderBuilder
	xmlBuilder     *builder.XMLBuilder
}

// NewConverterFacade crea una nueva instancia del facade
func NewConverterFacade() *ConverterFacade {
	return &ConverterFacade{
		headersBuilder: builder.NewHeadersBuilder(),
		orderBuilder:   builder.NewDistributionOrderBuilder(),
		xmlBuilder:     builder.NewXMLBuilder(),
	}
}

// ConversionRequest representa la entrada del facade
type ConversionRequest struct {
	Headers map[string]interface{} `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

// ConversionResult representa la salida del facade
type ConversionResult struct {
	Base64Gzip  string            `json:"base64_gzip"`
	XMLPath     string            `json:"xml_path"`
	TXTPath     string            `json:"txt_path"`
	OrderID     string            `json:"order_id"`
	HeadersInfo map[string]string `json:"headers_info,omitempty"`
}

// Convert realiza toda la operación de conversión de forma simplificada
func (f *ConverterFacade) Convert(request *ConversionRequest) (*ConversionResult, error) {
	// 1. Procesamos headers
	var headers *model.Headers
	var headersInfo map[string]string
	if request.Headers != nil {
		headers = f.headersBuilder.SetAllHeaders(request.Headers).Build()
		headersInfo = map[string]string{
			"domain":     headers.Domain,
			"event_type": headers.EventType,
			"entity_id":  headers.EntityId,
			"commerce":   headers.Commerce,
			"country":    headers.Country,
		}
	}

	// 2. Validamos body
	if request.Body == nil {
		return nil, fmt.Errorf("el campo 'body' es requerido")
	}

	// 3. Construimos DistributionOrder usando el patrón Builder
	order := f.orderBuilder.
		SetBasicInfo(request.Body).
		SetTimeInfo(request.Body).
		SetLabelInfo(request.Body).
		SetFacilityInfo(request.Body).
		SetLineItems(builder.GetArrayFromMap(request.Body, "lineItem")).
		Build()

	// 4. Validamos OrderID
	if order.DistributionOrderId == "" {
		return nil, fmt.Errorf("no se pudo obtener 'distributionOrderId' del body")
	}

	// 5. Convertimos, comprimimos y guardamos datos a XML, usando XMLBuilder
	f.xmlBuilder.SetXMLContent(order)

	// Guardamos XML
	xmlPath, err := f.xmlBuilder.SaveAsXML(fmt.Sprintf("%s.xml", order.DistributionOrderId))
	if err != nil {
		return nil, fmt.Errorf("error guardando XML: %v", err)
	}

	// Comprimimos a GZIP
	gzipped, err := f.xmlBuilder.CompressToGZIP()
	if err != nil {
		return nil, fmt.Errorf("error comprimiendo con GZIP: %v", err)
	}

	// Codificamos en Base64
	base64Encoded := f.xmlBuilder.EncodeToBase64(gzipped)

	// Guardamos TXT
	txtPath, err := f.xmlBuilder.SaveAsTXT(base64Encoded, fmt.Sprintf("%s.txt", order.DistributionOrderId))
	if err != nil {
		return nil, fmt.Errorf("error guardando TXT: %v", err)
	}

	// Retornamos resultado
	return &ConversionResult{
		Base64Gzip:  base64Encoded,
		XMLPath:     xmlPath,
		TXTPath:     txtPath,
		OrderID:     order.DistributionOrderId,
		HeadersInfo: headersInfo,
	}, nil
}
