package phong

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	defaultRegion = "JP"
	lang          = "en"
)

var (
	reg = defaultRegion
	num string
)

func setConfig(c *cli.Context) {
	num = c.Args().First()

	if len(num) == 0 {
		cli.ShowAppHelp(c)
		os.Exit(exitErr)
	}

	if c.NArg() >= 2 {
		err := fmt.Errorf("Pass one phone number as an argument")
		setError(err)
	}

	setRegion(c.String("region"))
}

func setRegion(r string) {
	if len(r) == 0 {
		return
	}

	validateRegion(r)
	reg = r
}

func validateRegion(r string) {
	if countryList[r] == nil {
		err := fmt.Errorf("Passed country code %q not found. Check the list of country code. For details see help", r)
		setError(err)
	}
}
