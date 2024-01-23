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
	// TODO: next steps:
	// temperature - random between 18 and 25
	// humidity - random between 0 and 1
	// pressure - random between 1000 and 1040
	// datetime - always `now`
	// id - constant, random string created on startup
	timestamp, _ := time.Parse(time.RFC3339, "2024-01-04T16:27:40Z")
	return TemperatureStats{15, 0.6, 1000, timestamp, "1"}
}

func JsonizeTemperatureStats(stats TemperatureStats) ([]byte, error) {
	return json.Marshal(stats)
}
