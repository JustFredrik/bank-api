package initializer

import (
	"encoding/xml"
	"io"
	"os"

	"github.com/justfredrik/bank-api/internal/camt053"
)

func LoadLocalData(path string) (camt053 camt053.Document, err error) {

	xmlFile, err := os.Open(path)
	if err != nil {
		return camt053, err
	}
	defer xmlFile.Close()

	byteData, _ := io.ReadAll(xmlFile)

	if err := xml.Unmarshal(byteData, &camt053); err != nil {
		return camt053, err
	}

	return camt053, err
}

var Bob = 45
