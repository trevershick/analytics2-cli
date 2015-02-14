package info

import (
	"bytes"
	"strings"
	"testing"
	"github.com/trevershick/analytics2-cli/a2m/test"
	"github.com/trevershick/analytics2-cli/a2m/config"
)

func Test_getWorkspaceInfoUrl(t *testing.T) {
	config := config.Configuration{}
	config.BaseUrl = "http://localhost:1000/xxx"

	url := getWorkspaceInfoUrl(&config, 103)
	test.AssertEquals(t, "http://localhost:1000/xxx/info/workspace/103", url)
}


func Test_showBasicInfoValue_BadFieldValue(t *testing.T) {
	cfg := config.Configuration{}
	cfg.BaseUrl = "http://localhost:1000/xxx"

	myLoader, passedInRestArgs := test.FakeRestLoader("{}")


	var myWriter bytes.Buffer

	args := &infoArgs{
		config: &cfg,
		workspaceId: 7,
		fieldName: "xxx",
		loader: myLoader,
		writer: &myWriter,
	}

	showBasicInfoValue(args)
	test.AssertEquals(t, "<invalid Value>", strings.TrimSpace(myWriter.String()))
	test.AssertEquals(t, "http://localhost:1000/xxx/info/workspace/7", passedInRestArgs.Url)
}



func Test_showBasicInfoValue_HappyPath(t *testing.T) {
	cfg := config.Configuration{}
	cfg.BaseUrl = "http://localhost:1000/xxx"

	myLoader, passedInRestArgs := test.FakeRestLoader(`{"Subscription":99}`)

	var myWriter bytes.Buffer

	args := &infoArgs{
		config: &cfg,
		workspaceId: 9,
		fieldName: "Subscription",
		loader: myLoader,
		writer: &myWriter,
	}

	showBasicInfoValue(args)
	test.AssertEquals(t, "99", strings.TrimSpace(myWriter.String()))
	test.AssertEquals(t, "http://localhost:1000/xxx/info/workspace/9", passedInRestArgs.Url)
}
