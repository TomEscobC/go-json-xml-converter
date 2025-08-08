package handler

import (
	"net/http"

	"github.com/TomEscobC/go-json-xml-converter/internal/service"
	"github.com/gin-gonic/gin"
)

// ConversionHandler maneja las solicitudes HTTP para la conversión.
type ConversionHandler struct {
	service *service.ConverterService
}

// NewConversionHandler crea una nueva instancia de ConversionHandler.
func NewConversionHandler(service *service.ConverterService) *ConversionHandler {
	return &ConversionHandler{service: service}
}

// conversionRequest representa la estructura de la solicitud de conversión.
// Esta definición es interna del handler para no exponer detalles del service/facade.
type conversionRequest struct {
	Headers map[string]interface{} `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

// Convert maneja la ruta POST /convert.
func (h *ConversionHandler) Convert(c *gin.Context) {
	var request conversionRequest

	// 1. Bind JSON
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	// 2. Convertimos a tipo del service
	serviceRequest := &service.ConversionRequest{
		Headers: request.Headers,
		Body:    request.Body,
	}

	// Llamamos al servicio
	result, err := h.service.Convert(serviceRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 4. Retornamos resultado
	c.JSON(http.StatusOK, result)
}
