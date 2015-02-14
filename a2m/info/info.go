package info

import (
	"github.com/codegangsta/cli"
)


func InfoCommands() []cli.Command {

	c := []cli.Command {
		{
			Name:"revision-count",
			Usage: "Show the number of revisions in a workspace's queue",
			Action: showBasicInfoValue("RevisionsInQueue"),
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
			Action: showBasicInfoValue("Halted"),
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

