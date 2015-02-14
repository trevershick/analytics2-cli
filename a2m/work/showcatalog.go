package work

import (
	"fmt"
	"os"
	"io"
	"net/url"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/config"
	"github.com/trevershick/analytics2-cli/a2m/rest"
)

type showCatalogArgs struct {
	config *config.Configuration
	loader rest.Loader
	writer io.Writer
}

func showCatalogCommand() cli.Command {
	return cli.Command {
			Name:"catalog",
			Usage: "Show the work catalog",
			Action: func(c *cli.Context) {
				args := showCatalogArgs{
					config: config.GetConfigurationOrPanic(c),
					loader: rest.ExecuteAndExtractJsonObject,
					writer: os.Stdout,
				}
				showCatalog(&args)
			},
		}
}

func getShowCatalogUrl(c *config.Configuration) string {
	return c.FullUrl("/management/work/catalog")
}

func showCatalog(args *showCatalogArgs) {

	showCatalogUrl := getShowCatalogUrl(args.config)
	params := url.Values{}

	var obj map[string]interface {}

	restArgs := &rest.RestArgs{
		Config: args.config,
		Url: showCatalogUrl,
		Params: params,
		ResponseData: &obj,
	}

	err := args.loader(restArgs)
	if err != nil {
		fmt.Fprintf(args.writer, "%s", err)
		os.Exit(1)
	}

	var available []interface {}
	available = obj["available"].([]interface{} )

	fmt.Fprintf(args.writer, "\nAvailable Tasks in the Work Catalog @ %s", showCatalogUrl, url.Values{})
	fmt.Fprintf(args.writer, "\n=============================================================================")
	for i := range available {
		workTask := available[i].(map[string]interface{})

		name := workTask["name"].(string)
		desc := workTask["description"].(string)
		fmt.Fprintf(args.writer, "\n%-40s : %s", name,desc )
	}
}
