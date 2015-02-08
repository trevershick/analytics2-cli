package rest

import (
	"fmt"
	"strings"
	"net/http"
	"net/url"
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


func ExecuteAndExtractJson(config *config.Configuration, url string, params url.Values) (map[string]interface{}, error) {
	client := &http.Client{}

	// move Urls to a method on the configuration object?

	fullUrl := []string {url, params.Encode()}

	req, err := http.NewRequest("GET", strings.Join(fullUrl, "?"), nil)
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

	var dat map[string]interface{}
	if err := json.Unmarshal(contents, &dat); err != nil {
		return nil, err
	}
	return dat, nil
}
