package main

import (
	"fmt"
	"log"

	"github.com/TomEscobC/go-json-xml-converter/internal/handler"
	"github.com/TomEscobC/go-json-xml-converter/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Inicializamos servicios
	converterService := service.NewConverterService()

	// Inicializamos handlers
	conversionHandler := handler.NewConversionHandler(converterService)

	// Configuramos router
	r := gin.Default()

	// Rutas
	r.POST("/convert", conversionHandler.Convert)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Go JSON to XML Converter is running"})
	})

	// Iniciamos servidor
	fmt.Println("Servidor iniciado en http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
