package utils

import (
	"encoding/json"
	"os"

	"transit-backend/internal/transport/model"
)

func LoadTransportData(path string) (*model.Data, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data model.Data
	err = json.Unmarshal(file, &data)
	return &data, err
}
