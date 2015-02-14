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


type Non200ResponseCode struct {
	code int
}

func (e Non200ResponseCode) Error() string {
	return fmt.Sprintf("Error %d has occurred", e.code)
}


func ExecuteAndExtractJsonObject(config *config.Configuration, url string, params url.Values, dat interface{}) (error) {
	client := &http.Client{}

	// move Urls to a method on the configuration object?

	fullUrl := []string {url, params.Encode()}

	req, err := http.NewRequest("GET", strings.Join(fullUrl, "?"), nil)
	req.SetBasicAuth(config.UserName, config.Password)
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

	if err := json.Unmarshal(contents, &dat); err != nil {
		return err
	}
	return nil
}

func ExecuteAndExtractJson(config *config.Configuration, url string, params url.Values) (map[string]interface{}, error) {
	var dat map[string]interface{}
	err := ExecuteAndExtractJsonObject(config, url, params, &dat)
	return dat, err
}
