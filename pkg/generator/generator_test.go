package generator

import (
	"encoding/json"
	"testing"
)

// Plan: refactor test, so that it checks that the TemperatureStats are within the expected bounds;
// write an assertion function that check specifically this type of objects; use throughout.
func TestGenerate(t *testing.T) {
	t.Run("should generate random data within the expected range", func(t *testing.T) {
		result := Generate()
		assertStatsInExpectedRange(t, result)
	})
}

func TestJsonizeTemperatureStats(t *testing.T) {
	t.Run("should unmarshal json correctly", func(t *testing.T) {
		result, err := JsonizeTemperatureStats(Generate())
		if err != nil {
			t.Errorf("failed getting generated data: %s", err)
		}

		unmarshalled := TemperatureStats{}
		err = json.Unmarshal(result, &unmarshalled)
		if err != nil {
			t.Errorf("failed to unmarshal temperature stats: %s", err)
		}
	})
}

func assertStatsInExpectedRange(t testing.TB, stats TemperatureStats) {
	t.Helper()

	if stats.Humidity < 0 || stats.Humidity > 1.0 {
		t.Errorf("Humidity should be between 0 and 1; got %f", stats.Humidity)
	}
	if stats.Temperature > 25 || stats.Temperature < 15 {
		t.Errorf("Temperature should be between 15 and 25; got %d", stats.Temperature)
	}
	if stats.Pressure > 1040 || stats.Pressure < 1000 {
		t.Errorf("Pressure should be between 1000 and 1040; got %d", stats.Temperature)
	}
}
