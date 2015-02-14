package info

import (
	"fmt"
	"os"
	"strconv"
	"net/url"
	"strings"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/trevershick/analytics2-cli/a2m/config"
)

func ShowRevisionCountCommand() cli.Command {
	return cli.Command {
		Name:"revision-count",
		Usage: "Show the number of revisions in a workspace's queue",
		Action: ShowRevisionCount,
		Flags: []cli.Flag {
			cli.IntFlag {
				Name: "workspace, w",
				Usage: "The workspace Id",
			},
		},
	}
}

type WorkspaceInfo struct {
	revisionsInQueue uint64
}

func getWorkspaceInfoUrl(c *cli.Context, config *config.Configuration) string {
	paths := []string {"/info/workspace", strconv.Itoa(c.Int("workspace"))}
	return config.FullUrl(strings.Join(paths, "/"))
}

func ShowRevisionCount(c *cli.Context) {
	config, err := config.GetConfiguration(c)
	if err != nil {
		panic(err)
	}

	params := url.Values{}

	if c.Bool("recent") {
		params.Set("active","false")
	}
	var workspaceInfo WorkspaceInfo
	err = rest.ExecuteAndExtractJsonObject(config, getWorkspaceInfoUrl(c, config), params, &workspaceInfo)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}


	fmt.Printf("%d", workspaceInfo.revisionsInQueue)
}
