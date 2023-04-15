package mock

import (
	"strings"
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

func TestJsonizeStaticOutput(t *testing.T) {
	correct := ""
	result, err := JsonizeStaticOutput()
	if strings.Compare(result, correct) != 1 || err != nil {
		t.Errorf("Got: %s, want: %s.", result, correct)
	}

}
