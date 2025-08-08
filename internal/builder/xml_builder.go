package builder

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// XMLBuilder implementa el patrón Builder para operaciones XML.
type XMLBuilder struct {
	xmlContent []byte
}

// NewXMLBuilder crea una nueva instancia de XMLBuilder.
func NewXMLBuilder() *XMLBuilder {
	return &XMLBuilder{}
}

// SetXMLContent establece el contenido XML a partir de una estructura.
func (b *XMLBuilder) SetXMLContent(data interface{}) *XMLBuilder {
	xmlBytes, err := xml.MarshalIndent(data, "", "  ")
	if err != nil {
		b.xmlContent = nil
		return b
	}

	b.xmlContent = []byte(xml.Header + string(xmlBytes))
	return b
}

// SaveAsXML guarda el contenido como archivo XML.
// filename debe ser solo el nombre del archivo, sin la carpeta 'output'.
func (b *XMLBuilder) SaveAsXML(filename string) (string, error) {
	if b.xmlContent == nil {
		return "", fmt.Errorf("no XML content to save")
	}

	fullPath := filepath.Join("output", filename)
	if err := b.saveToFile(fullPath, b.xmlContent); err != nil {
		return "", err
	}

	return fullPath, nil
}

// CompressToGZIP comprime el contenido en GZIP.
func (b *XMLBuilder) CompressToGZIP() ([]byte, error) {
	if b.xmlContent == nil {
		return nil, fmt.Errorf("no XML content to compress")
	}

	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(b.xmlContent); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// EncodeToBase64 codifica los datos en Base64.
func (b *XMLBuilder) EncodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// SaveAsTXT guarda el contenido como archivo TXT.
// content es el string a guardar.
// filename debe ser solo el nombre del archivo, sin la carpeta 'output'.
func (b *XMLBuilder) SaveAsTXT(content, filename string) (string, error) {
	fullPath := filepath.Join("output", filename)
	if err := b.saveToFile(fullPath, []byte(content)); err != nil {
		return "", err
	}

	return fullPath, nil
}

// saveToFile guarda datos en un archivo, creando directorios si es necesario.
func (b *XMLBuilder) saveToFile(filename string, data []byte) error {
	// Creamos directorio si no existe
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// GenerateTimestamp genera un timestamp para nombres de archivo.
func (b *XMLBuilder) GenerateTimestamp() string {
	return time.Now().Format("20060102_150405")
}
