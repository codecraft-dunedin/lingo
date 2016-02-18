package app

import (
	"fmt"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/lingo-reviews/lingo/commands"
	"github.com/lingo-reviews/lingo/commands/common"
)

// TODO(waigani) have a global state that tenets can share. An issue may be
// found based on the combination of tenets.

func New() *cli.App {
	setCommandHelpTemplate()
	app := cli.NewApp()
	app.Name = "lingo"
	app.Usage = "A DevOps Tool for Software Engineering. This is an ALPHA release (here be BUGS)."
	app.Before = commands.BeforeCMD
	app.Commands = globalCommands
	app.Flags = common.GlobalOptions
	app.CommandNotFound = commands.TenetCMD
	app.EnableBashCompletion = true

	fmt.println("Hello World")
	
	return app
}

func setCommandHelpTemplate(args ...string) {
	var argStr string
	if len(args) == 0 {
		argStr = "[arguments...]"
	} else {
		argStr = "<" + strings.Join(args, "> <") + ">"
	}
	cli.CommandHelpTemplate = fmt.Sprintf(`
NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   lingo {{.Name}}{{if .Flags}} [options]{{end}} %s{{if .Description}}

EXAMPLES:
   {{.Description}}{{end}}{{if .Flags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{ end }}
`[1:], argStr)
}
