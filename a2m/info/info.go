package info

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/config"
	"github.com/trevershick/analytics2-cli/a2m/rest"
)


func InfoCommands() []cli.Command {

	showBasicInfoAction := func(name string) func(c *cli.Context) {
		return func(c *cli.Context) {
			args := infoArgs{
				config: config.GetConfigurationOrPanic(c),
				fieldName : name,
				workspaceId : getWorkspaceId(c),
				loader: rest.ExecuteAndExtractJsonObject,
				writer: os.Stdout,
			}
			showBasicInfoValue(&args)
		}
	}


	c := []cli.Command {
		{
			Name:"revision-count",
			Usage: "Show the number of revisions in a workspace's queue",
			Action: showBasicInfoAction("RevisionsInQueue"),
			Flags: []cli.Flag {
				cli.IntFlag {
					Name: "workspace, w",
					Usage: "The workspace Id",
				},
			},
		},
		{
			Name:"halted",
			Usage: "Shows if a workspace is halted",
			Action: showBasicInfoAction("Halted"),
			Flags: []cli.Flag {
				cli.IntFlag {
					Name: "workspace, w",
					Usage: "The workspace Id",
				},
			},
		},
		{
			Name: "collections",
			Usage: "Show basic workspace collection information",
			Action: func (c *cli.Context) {
				args := showCollectionArgs {
					config: config.GetConfigurationOrPanic(c),
					workspaceId: getWorkspaceId(c),
					loader: rest.ExecuteAndExtractJsonObject,
					writer: os.Stdout,
				}
				showCollectionInformation(&args)
			},
			Flags: []cli.Flag {
				cli.IntFlag {
					Name: "workspace, w",
					Usage: "The workspace Id",
				},
			},
		},
		{
			Name: "halted-workspaces",
			Usage: "Show halted workspaces",
			Action: func (c *cli.Context) {
				args := showHaltedArgs {
					config: config.GetConfigurationOrPanic(c),
					loader: rest.ExecuteAndExtractJsonObject,
					writer: os.Stdout,
				}
				showHalted(&args)
			},
		},
	}
	return c
}

func getWorkspaceId(c *cli.Context) int {
	return c.Int("workspace")
}
