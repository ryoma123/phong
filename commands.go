package phong

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

// Flags variable
var Flags = []cli.Flag{
	flagName,
}

// Commands variable
var Commands = []*cli.Command{
	commandList,
}

var usages = map[string]string{
	"list": "| phong l",
}

var flagName = &cli.StringFlag{
	Name:        "region, r",
	Aliases:     []string{"r"},
	Value:       "JP",
	Usage:       "Pass a `<country code>`",
	DefaultText: "JP",
}

var commandList = &cli.Command{
	Name:    "list",
	Aliases: []string{"l"},
	Usage:   "Show the list of country code (alpha-2)",
	Action:  doList,
}

func init() {
	usages := setUsages()
	cli.CommandHelpTemplate = `NAME:
    {{.Name}} - {{.Usage}}
USAGE:
    phong {{.Name}} ` + usages + `{{ "\n\n" }}`
}

func setUsages() string {
	s := "{{if false}}"
	for _, command := range append(Commands) {
		s = s + fmt.Sprintf("{{else if (eq .Name %q)}}%s", command.Name, usages[command.Name])
	}
	return s + "{{end}}"
}

func doList(c *cli.Context) error {
	if c.NArg() != 0 {
		cli.ShowCommandHelp(c, "list")
		os.Exit(exitErr)
	}

	confirm(len(countryList))

	for k, v := range countryList {
		fmt.Printf("%s  %s -- %s\n", v.flag, k, v.name)
	}
	return nil
}

func confirm(n int) {
	fmt.Printf("Display all %d lines? (Y or n)", n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()

	if strings.ToLower(t) == "y" || len(t) == 0 {
		return
	}
	os.Exit(exitErr)
}
