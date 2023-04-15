package mock

import (
	"encoding/json"
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
	return json.Marshal(GetStaticOutput())
}
