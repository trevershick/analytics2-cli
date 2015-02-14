package info

import (
	"io"
	"fmt"
	"os"
	"net/url"
	"github.com/trevershick/analytics2-cli/a2m/config"
	"github.com/trevershick/analytics2-cli/a2m/rest"
)

type showHaltedArgs struct {
	config *config.Configuration
	loader rest.Loader
	writer io.Writer
}

func getHaltedWorkspaceUrl(config *config.Configuration) string {
	return config.FullUrl("/info/haltedWorkspaces")
}

func showHalted(args *showHaltedArgs) {
	wi := []HaltedWorkspaceInfo{}

	restArgs := &rest.RestArgs{
		Config: args.config,
		Url: getHaltedWorkspaceUrl(args.config),
		Params: url.Values{},
		ResponseData: &wi,
	}
	err := args.loader(restArgs)

	if err != nil {
		fmt.Fprintf(args.writer, "%s", err)
		os.Exit(1)
	}

	fmt.Fprintf(args.writer, "\nHalted Workspaces")
	fmt.Fprintf(args.writer, "\n=============================================================================")
	for _, c := range wi {
		r := c.Data.Reason
		w := uint64(c.WorkspaceOid)
		s := uint64(c.SubscriptionId)
		h := c.HealthCheckShouldFail
		t := c.Timestamp

		fmt.Fprintf(args.writer, "\n%-15d %-15d %5t %-15s %s", w, s, h, t, r)
	}
	fmt.Fprintf(args.writer, "\n\n")
}
