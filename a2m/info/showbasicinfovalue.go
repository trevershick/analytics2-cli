package info

import (
	"io"
	"fmt"
	"os"
	"strconv"
	"reflect"
	"net/url"
	"strings"
	"github.com/trevershick/analytics2-cli/a2m/rest"
	"github.com/trevershick/analytics2-cli/a2m/config"
)

type infoArgs struct {
	config *config.Configuration
	fieldName string
	workspaceId int
	responseData interface{}
	loader rest.Loader
	writer io.Writer
}

func getWorkspaceInfoUrl(config *config.Configuration, workspaceId int) string {
	paths := []string {"/info/workspace", strconv.Itoa(workspaceId)}
	return config.FullUrl(strings.Join(paths, "/"))
}

func showBasicInfoValue(args *infoArgs) {
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
	// gracias - http://stackoverflow.com/questions/18930910/golang-access-struct-property-by-name
	r := reflect.ValueOf(wi)
	f := reflect.Indirect(r).FieldByName(args.fieldName)

	switch f.Kind() {
	default:
		fmt.Fprintf(args.writer, "%v\n", f)
	case reflect.Float64:
		fmt.Fprintf(args.writer, "%0.0f\n", f.Float())
	case reflect.String:
		fmt.Fprintf(args.writer, "%s\n", f.String())
	case reflect.Bool:
		fmt.Fprintf(args.writer, "%t\n", f.Bool())
	}
}

