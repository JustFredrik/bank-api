package initializer

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/justfredrik/bank-api/internal/camt053"
)

func LoadLocalData(path string) (camt053.Document, error) {

	var doc camt053.Document
	xmlFile, err := os.Open(path)
	if err != nil {
		return doc, err
	}
	defer xmlFile.Close()

	byteData, _ := io.ReadAll(xmlFile)

	if err := xml.Unmarshal(byteData, &doc); err != nil {
		return doc, err
	}

	return doc, nil
}
