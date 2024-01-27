package generator

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

var id string = uuid.New().String()

type TemperatureStats struct {
	Temperature int       `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Pressure    int       `json:"pressure"`
	Datetime    time.Time `json:"datetime"`
	Id          string    `json:"id"`
}

func Generate() TemperatureStats {
	// id - constant, random string created on startup
	return TemperatureStats{rand.Intn(10) + 15, rand.Float32(), rand.Intn(40) + 1000, time.Now(), id}
}

func JsonizeTemperatureStats(stats TemperatureStats) ([]byte, error) {
	return json.Marshal(stats)
}
