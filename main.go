package main

import (
	"log"
	"os"

	"github.com/LintaoAmons/string-go/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:            "nric",
		Description:     "A tool to generate testing Singapore NRIC number, and it will automaticlly copy into you clipboard",
		HideHelpCommand: true,
		UsageText:       "nric [count] -- count is set to 1 if not passed",
		Usage:           "nric [count]",
		Action:          commands.ToCamelCaseCommand.Action,
		Commands: []*cli.Command{
			commands.ToCamelCaseCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
