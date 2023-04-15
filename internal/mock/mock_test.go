package mock

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStaticOutput(t *testing.T) {

	correct := MockTemperatureStats{15, 0.6, 1000}

	result := GetStaticOutput()
	if !cmp.Equal(result, correct) {
		t.Errorf("Got: %v, want: %v.", result, correct)
	}
}

func TestJsonizeStaticOutput(t *testing.T) {
	correct := []byte("{\"temperature\":15,\"humidity\":0.6,\"pressure\":1000}")
	result, err := JsonizeStaticOutput()
	if err != nil {
		t.Errorf("Error during testing: %s", err)
	}
	if !bytes.Equal(result, correct) {
		t.Errorf("Got: %s, want: %s.", result, correct)
	}

}
