package specifications

import (
	"bytes"
	"testing"
)

// DataGenerator mimics the domain interface for generating data for use in tests.
// TODO: consider using the domain interface directly and removing this one.
type DataGenerator interface {
	GenerateData() ([]byte, error)
}

// GenerateDataSpecification tests the domain specification of the data generator.
func GenerateDataSpecification(t testing.TB, dg DataGenerator) {
	got, err := dg.GenerateData()

	if err != nil {
		t.Fatalf("could not generate data: %s", err)
	}

	want := []byte("{\"temperature\":15,\"humidity\":0.6,\"pressure\":1000,\"datetime\":\"0001-01-01T00:00:00Z\",\"id\":\"1\"}") // TODO validate that output is parsable as TemperatureStats instead of hassling with strings
	if !bytes.Equal(got, want) {
		t.Errorf("wrong output, got %v, want %v", got, want)
	}
}
