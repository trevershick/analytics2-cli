package info

import (
	"io"
	"fmt"
	"os"
	"net/url"
	"github.com/trevershick/analytics2-cli/a2m/config"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/pivotal-golang/bytefmt"
)

type showCollectionArgs struct {
	config *config.Configuration
	workspaceId int
	loader rest.Loader
	writer io.Writer
}

func showCollectionInformation(args *showCollectionArgs) {
	wi := WorkspaceInfo{}

	restArgs := &rest.RestArgs{
		Config: args.config,
		Url: getWorkspaceInfoUrl(args.config, args.workspaceId),
		Params: url.Values{},
		ResponseData: &wi,
	}
	err := args.loader(restArgs)

	if err != nil {
		fmt.Fprintf(args.writer, "%s", err)
		os.Exit(1)
	}

	fmt.Fprintf(args.writer, "\nCollections for Workspace %d", args.workspaceId)
	fmt.Fprintf(args.writer, "\n=============================================================================")
	for _, c := range wi.Collections {
		n := c.Name
		s := c.TotalStorageSize
		fmt.Fprintf(args.writer, "\n%-40s %-15d %-10s", n, uint64(s), bytefmt.ByteSize(uint64(s)))
	}
	fmt.Fprintf(args.writer, "\n\n")
}
