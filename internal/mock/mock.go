package mock

import (
	"encoding/json"
	"fmt"
)

type MockTemperatureStats struct {
	Temperature int     `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    int     `json:"pressure"`
}

func GetStaticOutput() MockTemperatureStats {
	return MockTemperatureStats{15, 0.6, 1000}
}

func JsonizeStaticOutput() ([]byte, error) {
	jsonOutput, err := json.Marshal(GetStaticOutput())
	if err != nil {
		fmt.Printf("Error during jsonizing: %s", err)
		return nil, err
	}
	return jsonOutput, nil
}
