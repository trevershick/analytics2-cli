package rest

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"github.com/trevershick/analytics2-cli/a2m/config"
)

type Loader (func(*RestArgs) error)

type RestArgs struct {
	Config *config.Configuration
	Url string
	Params url.Values
	ResponseData interface{}
}

type Non200ResponseCode struct {
	code int
}

func (e Non200ResponseCode) Error() string {
	return fmt.Sprintf("Error %d has occurred", e.code)
}

func ExecuteAndExtractJsonObject(args *RestArgs) (error) {
	client := &http.Client{}

	// move Urls to a method on the configuration object?
	// fmt.Println("Getting data %s", url)

	fullUrl := []string {args.Url, args.Params.Encode()}

	req, err := http.NewRequest("GET", strings.Join(fullUrl, "?"), nil)
	req.SetBasicAuth(args.Config.UserName, args.Config.Password)
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return Non200ResponseCode{code:resp.StatusCode}
	}

	if err := json.Unmarshal(contents, args.ResponseData); err != nil {
		return err
	}
	return nil
}
