package work

import (
	"fmt"
	"os"
	"net/url"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/trevershick/analytics2-cli/a2m/config"
)

func ShowCatalogCommand() cli.Command {
	return cli.Command {
			Name:"catalog",
			Usage: "Show the work catalog",
			Action: ShowCatalog,
		}
}

func getShowCatalogUrl(c *config.Configuration) string {
	return c.FullUrl("/management/work/catalog")
}

func ShowCatalog(c *cli.Context) {

	config, err := config.GetConfiguration(c)
	if err != nil {
		panic(err)
	}

	obj, err := rest.ExecuteAndExtractJson(config, getShowCatalogUrl(config), url.Values{})
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println("Yay, got results -> ", obj)

	var available []interface {}
	available = obj["available"].([]interface{} )

	fmt.Printf("\nAvailable Tasks in the Work Catalog @ %s", getShowCatalogUrl(config), url.Values{})
	fmt.Printf("\n=============================================================================")
	for i := range available {
		workTask := available[i].(map[string]interface{})

		name := workTask["name"].(string)
		desc := workTask["description"].(string)
		fmt.Printf("\n%-40s : %s", name,desc )
	}
}
