package work

import (
	"fmt"
	"os"
	"io"
	"net/url"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/trevershick/analytics2-cli/a2m/config"
)

type showTasksArgs struct {
	config *config.Configuration
	recent bool
	loader rest.Loader
	writer io.Writer
}

func showTasksCommand() cli.Command {
	return cli.Command {
		Name:"tasks",
		Usage: "Show the tasks in-flight",
		Action: func(c *cli.Context) {
			args := &showTasksArgs{
				config: config.GetConfigurationOrPanic(c),
				recent: c.Bool("recent"),
				loader: rest.ExecuteAndExtractJsonObject,
				writer: os.Stdout,
			}
			showTasks(args)
		},
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

func showTasks(args *showTasksArgs) {

	params := url.Values{}

	if args.recent {
		params.Set("active","false")
	}
	tasksUrl := getTasksUrl(args.config)


	var obj map[string]interface{}

	restArgs := &rest.RestArgs{
		Config: args.config,
		Url: tasksUrl,
		Params: params,
		ResponseData: &obj,
	}

	err := args.loader(restArgs)
	if err != nil {
		fmt.Fprintf(args.writer, "%s", err)
		os.Exit(1)
	}

	var tasks []interface {}
	tasks = obj["tasks"].([]interface{} )

	fmt.Fprintf(args.writer, "\nActive Tasks @ %s", tasksUrl)
	fmt.Fprintf(args.writer, "\n=============================================================================")
	for i := range tasks {
		t := tasks[i].(map[string]interface{})

		name := t["taskKey"].(string)
		status := t["status"].(string)
		start := t["startDate"].(string)
		end := t["endDate"].(string)
		fmt.Fprintf(args.writer, "\n%-40s %-10s %-15s %-15s", name,status,start,end )
	}
	fmt.Fprintf(args.writer, "\n\n")
}
