package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/work"
	"github.com/trevershick/analytics2-cli/a2m/info"
)

func main() {

	// fmt.Printf("%s",config)
	app := cli.NewApp()
	app.Name = "a2m"
	app.Author = "Trever Shick"
	app.Email = "tshick@rallydev.com"
	app.Usage = "A2 Management CLI"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "base, b",
			Value: "http://localhost:9201/analytics-etl",
			Usage: "Base URL to connect to",
			EnvVar: "A2M_BASE_URL",
		},
		cli.StringFlag {
			Name: "user, u",
			Value: "nobody",
			Usage: "User Name to connect to the web service",
			EnvVar: "A2M_USER",
		},
		cli.StringFlag {
			Name: "pass, p",
			Value: "nothing",
			Usage: "Password to use for authentication",
			EnvVar: "A2M_PASSWORD",
		},
		cli.BoolFlag {
			Name: "save, s",
			Usage: "Save supplied arguments to ~/.a2mrc",
		},
	}
	commands := []cli.Command{}

	// put into a loop of functions
	x := work.WorkCommands()
	commands = append(commands, x...)

	x = info.InfoCommands()
	commands = append(commands, x...)

	app.Commands = commands
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
