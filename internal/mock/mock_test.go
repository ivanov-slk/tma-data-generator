package mock

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetStaticOutput(t *testing.T) {

	correct := MockTemperatureStats{15, 0.6, 1000}

	out, err := GetStaticOutput()
	if !cmp.Equal(out, correct) || err != nil {
		t.Fail()
	}
}
