package admin

import (
	"io"
	"fmt"
	"os"
	"strconv"
	"net/url"
	"github.com/trevershick/analytics2-cli/a2m/config"
	"github.com/trevershick/analytics2-cli/a2m/rest"
)

type unhaltArgs struct {
	config *config.Configuration
	workspaceId int
	loader rest.Loader
	writer io.Writer
}

func getUnhaltUrl(config *config.Configuration) string {
	return config.FullUrl("/admin/etlUnhaltWorkspace")
}

func unhalt(args *unhaltArgs) {
	params := url.Values{}
	params.Set("id", strconv.Itoa(args.workspaceId))

	restArgs := &rest.RestArgs{
		Config: args.config,
		Url: getUnhaltUrl(args.config),
		Params: params,
	}

	err := args.loader(restArgs)

	if err != nil {
		fmt.Fprintf(args.writer, "%s\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(args.writer, "%v\n", restArgs.ResponseData)
}
