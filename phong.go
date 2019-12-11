package phong

import (
	"fmt"

	"github.com/nyaruka/phonenumbers"
)

// Phong struct
type Phong struct {
	name     string
	flag     string
	location string
}

// New returns Phong
func New(n string) *Phong {
	p, err := phonenumbers.Parse(n, reg)
	if err != nil {
		err := fmt.Errorf(err.Error())
		setError(err)
	}

	r := phonenumbers.GetRegionCodeForNumber(p)
	c := newCountry(r)
	if c == nil {
		err := fmt.Errorf("Region not found")
		setError(err)
	}

	l, err := phonenumbers.GetGeocodingForNumber(p, "en")
	if err != nil {
		err := fmt.Errorf(err.Error())
		setError(err)
	}

	return &Phong{
		name:     c.name,
		flag:     c.flag,
		location: l,
	}
}

func (p Phong) output() {
	if len(p.location) == 0 {
		fmt.Printf("%s  %s\n", p.flag, p.name)
	} else {
		fmt.Printf("%s  %s -- %s\n", p.flag, p.name, p.location)
	}
}
