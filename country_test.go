package phong

import (
	"reflect"
	"testing"
)

func TestNewCountry(t *testing.T) {
	m := map[string]*country{
		"AF": &country{flag: "🇦🇫", name: "Afghanistan"},
		"ZW": &country{flag: "🇿🇼", name: "Zimbabwe"},
	}

	for k := range m {
		if !reflect.DeepEqual(newCountry(k), m[k]) {
			t.Errorf("Failed newCountry: %s", k)
		}
	}
}
