package phong

import (
	"github.com/urfave/cli/v2"
)

// RunCLI runs as cli
func RunCLI(c *cli.Context) error {
	setConfig(c)

	p := New(num)
	p.output()
	return nil
}
