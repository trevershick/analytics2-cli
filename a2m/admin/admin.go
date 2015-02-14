package admin

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/config"
	"github.com/trevershick/analytics2-cli/a2m/rest"
)


func AdminCommands() []cli.Command {


	c := []cli.Command {
		{
			Name: "halt",
			Usage: "Halt a workspace",
			Action: func (c *cli.Context) {
				args := haltArgs {
					workspaceId: getWorkspaceId(c),
					config: config.GetConfigurationOrPanic(c),
					loader: rest.ExecuteAndExtractPlainText,
					writer: os.Stdout,
				}
				halt(&args)
			},
			Flags: []cli.Flag {
				cli.IntFlag {
					Name: "workspace, w",
					Usage: "The workspace Id",
				},
			},
		},
	}
	return c
}

func getWorkspaceId(c *cli.Context) int {
	return c.Int("workspace")
}
