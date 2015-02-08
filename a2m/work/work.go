package work

import (
	"net/http"
	"github.com/codegangsta/cli"
	"fmt"
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"
	"github.com/RallySoftware/analytics2-cli/a2m/config"
)

type Non200ResponseCode struct {
	code int
}

func (e Non200ResponseCode) Error() string {
	return fmt.Sprintf("Error %d has occurred", e.code)
}

type JsonObject map[string]interface {

}

type Urls struct {
	base string
}
func GetUrls() (Urls, error) {
	return Urls {
		base : "http://localhost:9201/analytics-etl",
	}, nil
}

func (urls Urls) GetShowCatalogUrl() string {
	s := []string { urls.base, "/management/work/catalog" }
	return strings.Join(s, "")
}

func WorkCommands() []cli.Command {
	workCommands := []cli.Command {
		{
			Name:"catalog",
			Usage: "Show the work catalog",
			Action: showCatalog,
		},
	}
	return workCommands
}

func ExecuteAndExtractJson(url string) (JsonObject, error) {
	client := &http.Client{}

	config, err := config.GetConfiguration()
	if err != nil {
		return nil, err
	}

	// move Urls to a method on the configuration object?

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(config.UserName, config.Password)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, Non200ResponseCode{code:resp.StatusCode}
	}

	var dat JsonObject
	if err := json.Unmarshal(contents, &dat); err != nil {
		return nil, err
	}
	return dat, nil
}

func showCatalog(c *cli.Context) {
	println("Showing Catalog: ", c.Args().First())

	urls, _ := GetUrls()

	obj, err := ExecuteAndExtractJson(urls.GetShowCatalogUrl())
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Println("Yay, got results -> ", obj)

	var available []interface {}
	available = obj["available"].([]interface{} )

	fmt.Printf("\nAvailable Tasks in the Work Catalog @ %s", urls.base)
	fmt.Printf("\n=============================================================================")
	for i := range available {
		workTask := available[i].(map[string]interface{})

		name := workTask["name"].(string)
		desc := workTask["description"].(string)
		fmt.Printf("\n%-40s : %s", name,desc )
	}

}
