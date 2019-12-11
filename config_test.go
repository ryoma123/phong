package phong

import (
	"testing"
)

func TestSetRegion(t *testing.T) {
	if e := "JP"; reg != e {
		t.Errorf("Failed reg: %s", e)
	}

	setRegion("US")

	if e := "US"; reg != e {
		t.Errorf("Failed setRegion: %s", e)
	}
}
