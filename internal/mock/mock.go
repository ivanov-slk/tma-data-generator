package mock

type MockTemperatureStats struct {
	Temperature int     `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Pressure    int     `json:"pressure"`
}

func GetStaticOutput() (MockTemperatureStats, error) {
	return MockTemperatureStats{15, 0.6, 1000}, nil
}
