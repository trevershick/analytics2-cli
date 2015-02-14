package info

import (
	"fmt"
	"os"
	"net/url"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/trevershick/analytics2-cli/a2m/config"
	"github.com/pivotal-golang/bytefmt"
)

func showCollectionInformation(c *cli.Context) {
	config, err := config.GetConfiguration(c)
	if err != nil {
		panic(err)
	}

	wi := WorkspaceInfo{}
	err = rest.ExecuteAndExtractJsonObject(config, getWorkspaceInfoUrl(c, config), url.Values{}, &wi)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("\nCollections for Workspace %d", uint64(wi.Workspace))
	fmt.Printf("\n=============================================================================")
	for _, c := range wi.Collections {
		n := c.Name
		s := c.TotalStorageSize
		fmt.Printf("\n%-40s %-15d %-10s", n, uint64(s), bytefmt.ByteSize(uint64(s)))
	}
	fmt.Printf("\n\n")
}
