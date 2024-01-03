package generator

import (
	"encoding/json"
	"time"
)

type TemperatureStats struct {
	Temperature int       `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Pressure    int       `json:"pressure"`
	Datetime    time.Time `json:"datetime"`
	Id          string    `json:"id"`
}

func Generate() TemperatureStats {
	timestamp, _ := time.Parse(time.RFC3339, "2024-01-04T16:27:40Z00:00")
	return TemperatureStats{15, 0.6, 1000, timestamp, "1"}
}

func JsonizeTemperatureStats() ([]byte, error) {
	return json.Marshal(Generate())
}
