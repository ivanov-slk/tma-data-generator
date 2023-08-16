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

	want := []byte{1}
	if !bytes.Equal(got, want) {
		t.Errorf("wrong output, got %v, want %v", got, want)
	}
}
