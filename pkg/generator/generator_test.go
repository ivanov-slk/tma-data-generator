package generator

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2024-01-04T16:27:40Z")
	correct := TemperatureStats{15, 0.6, 1000, timestamp, "1"}

	result := Generate()
	if !reflect.DeepEqual(result, correct) { // for now shallow comparison is enough
		t.Errorf("Got: %v, want: %v.", result, correct)
	}
}

func TestJsonizeTemperatureStats(t *testing.T) {
	correct := []byte("{\"temperature\":15,\"humidity\":0.6,\"pressure\":1000,\"datetime\":\"2024-01-04T16:27:40Z\",\"id\":\"1\"}")
	result, err := JsonizeTemperatureStats(Generate())
	if err != nil {
		t.Errorf("Error during testing: %s", err)
	}
	if !bytes.Equal(result, correct) {
		t.Errorf("Got: %s, want: %s.", result, correct)
	}

}
