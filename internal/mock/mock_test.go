package mock

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStaticOutput(t *testing.T) {

	correct := MockTemperatureStats{15, 0.6, 1000}

	result, err := GetStaticOutput()
	if !cmp.Equal(result, correct) || err != nil {
		t.Errorf("Got: %v, want: %v.", result, correct)
	}
}
