package work

import (
	"fmt"
	"os"
	"net/url"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/trevershick/analytics2-cli/a2m/config"
)

func showTasksCommand() cli.Command {
	return cli.Command {
		Name:"tasks",
		Usage: "Show the tasks in-flight",
		Action: showTasks,
		Flags: []cli.Flag {
			cli.BoolFlag {
				Name: "recent, r",
				Usage: "Show recent tasks as well as active tasks",
			},
		},
	}
}

func getTasksUrl(c *config.Configuration) string {
	return c.FullUrl("/management/work/tasks")
}

func showTasks(c *cli.Context) {
	config, err := config.GetConfiguration(c)
	if err != nil {
		panic(err)
	}

	params := url.Values{}

	if c.Bool("recent") {
		params.Set("active","false")
	}
	obj, err := rest.ExecuteAndExtractJson(config, getTasksUrl(config), params)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println("Yay, got results -> ", obj)

	var tasks []interface {}
	tasks = obj["tasks"].([]interface{} )

	fmt.Printf("\nActive Tasks @ %s", getTasksUrl(config))
	fmt.Printf("\n=============================================================================")
	for i := range tasks {
		t := tasks[i].(map[string]interface{})

		name := t["taskKey"].(string)
		status := t["status"].(string)
		start := t["startDate"].(string)
		end := t["endDate"].(string)
		fmt.Printf("\n%-40s %-10s %-15s %-15s", name,status,start,end )
	}
	fmt.Printf("\n\n")
}
