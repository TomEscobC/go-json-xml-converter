package service

import (
	"github.com/TomEscobC/go-json-xml-converter/internal/facade"
)

type ConversionRequest struct {
	Headers map[string]interface{} `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

type ConversionResult struct {
	Base64Gzip  string            `json:"base64_gzip"`
	XMLPath     string            `json:"xml_path"`
	TXTPath     string            `json:"txt_path"`
	OrderID     string            `json:"order_id"`
	HeadersInfo map[string]string `json:"headers_info,omitempty"`
}

type ConverterService struct {
	converterFacade *facade.ConverterFacade
}

// NewConverterService crea una nueva instancia de ConverterService.
func NewConverterService() *ConverterService {
	return &ConverterService{
		converterFacade: facade.NewConverterFacade(),
	}
}

// Convert realiza la conversión usando el facade.
func (s *ConverterService) Convert(request *ConversionRequest) (*ConversionResult, error) {
	// Convertir el request del service al request del facade
	facadeRequest := &facade.ConversionRequest{
		Headers: request.Headers,
		Body:    request.Body,
	}

	// Llamar al facade
	facadeResult, err := s.converterFacade.Convert(facadeRequest)
	if err != nil {
		return nil, err
	}

	// Retornar el resultado (mismo tipo)
	return &ConversionResult{
		Base64Gzip:  facadeResult.Base64Gzip,
		XMLPath:     facadeResult.XMLPath,
		TXTPath:     facadeResult.TXTPath,
		OrderID:     facadeResult.OrderID,
		HeadersInfo: facadeResult.HeadersInfo,
	}, nil
}
