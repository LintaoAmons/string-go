package commands

import (
	"fmt"
	"strings"

	"io/ioutil"

	"github.com/LintaoAmons/string-go/functions"
	"github.com/atotto/clipboard"
	"github.com/urfave/cli/v2"
)

var ToCamelCaseCommand = &cli.Command{
	Name:        "to-camel-case",
	Description: "convert input to camel case",
	Action: func(ctx *cli.Context) error {
		// TODO wait to read user input
		input := ctx.Args().Get(0)
		output := functions.ToCamelCase(input)
		clipboard.WriteAll(output)
		fmt.Println(output)
		return nil
	},
	Subcommands: []*cli.Command{
		{
			Name: "file",
			Action: func(ctx *cli.Context) error {
				filepath := ctx.Args().Get(0)
				content := MustReadFile(filepath)
				lines := strings.Split(content, "\n")
				output := []string{}
				for _, line := range lines {
					output = append(output, functions.ToCamelCase(line))
				}
				outputString := strings.Join(output, "\n")
				clipboard.WriteAll(outputString)
				fmt.Println(outputString)
				return nil
			},
		},
	},
}

func MustReadFile(filepath string) string {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(b)
}
