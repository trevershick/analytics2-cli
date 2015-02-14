package info

import (
	"fmt"
	"os"
	"strconv"
	"reflect"
	"net/url"
	"strings"
	"github.com/codegangsta/cli"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/trevershick/analytics2-cli/a2m/config"
)


func getWorkspaceInfoUrl(c *cli.Context, config *config.Configuration) string {
	paths := []string {"/info/workspace", strconv.Itoa(c.Int("workspace"))}
	return config.FullUrl(strings.Join(paths, "/"))
}

func showBasicInfoValue(fieldName string) func(*cli.Context) {
	return func(c *cli.Context) {
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
		// gracias - http://stackoverflow.com/questions/18930910/golang-access-struct-property-by-name
		r := reflect.ValueOf(wi)
		f := reflect.Indirect(r).FieldByName(fieldName)

		switch f.Kind() {
		default:
			fmt.Printf("%v\n", f)
		case reflect.Float64:
			fmt.Printf("%0.0f\n", f.Float())
		case reflect.String:
			fmt.Printf("%s\n", f.String())
		case reflect.Bool:
			fmt.Printf("%t\n", f.Bool())
		}
	}
}

