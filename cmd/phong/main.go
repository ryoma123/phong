package main

import (
	"os"

	"github.com/ryoma123/phong"
	"github.com/urfave/cli/v2"
)

var version = "0.1.0"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := &cli.App{
		Name:    "phong",
		Version: version,
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "ryoma123",
				Email: "ryoma.ono.2661@gmail.com",
			},
		},
		Usage:    "CLI tool to inform the region where a phone number is from",
		Flags:    phong.Flags,
		Commands: phong.Commands,
		Action: func(c *cli.Context) error {
			phong.RunCLI(c)
			return nil
		},
	}
	return app
}
