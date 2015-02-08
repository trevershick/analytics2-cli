package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/RallySoftware/analytics2-cli/a2m/work"
	"github.com/RallySoftware/analytics2-cli/a2m/config"
)

func main() {

	_, err := config.GetConfiguration()
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%s",config)
	app := cli.NewApp()
	app.Name = "a2m"
	app.Usage = "fight the loneliness!"
	app.Commands = work.WorkCommands()
/*
	app.Commands = []cli.Command {
	{
		Name:"add",
		ShortName: "a",
		Usage: "add a task to the list",
		Action: func(c *cli.Context) {
			println("added task: ", c.Args().First())
		},
	},
	{
		Name:"complete",
		ShortName: "c",
		Usage: "complete a task on the list",
		Action: func(c *cli.Context) {
			println("completed task: ", c.Args().First())
		},
	},
	{
		Name:"template",
		ShortName: "r",
		Usage: "options for task templates",
		Subcommands: []cli.Command{
		{
			Name:"add",
			Usage: "add a new template",
			Action: func(c *cli.Context) {
				println("new task template: ", c.Args().First())
			},
		},
		{
			Name:"remove",
			Usage: "remove an existing template",
			Action: func(c *cli.Context) {
				println("removed task template: ", c.Args().First())
			},
		},
		},
	},
	}*/


	app.Run(os.Args)
}
