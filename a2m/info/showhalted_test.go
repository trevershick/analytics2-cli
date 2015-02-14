package info

import (
	"bytes"
	"testing"
	"github.com/trevershick/analytics2-cli/a2m/test"
	"github.com/trevershick/analytics2-cli/a2m/config"
)


func Test_showHalted(t *testing.T) {
	cfg := config.Configuration{}
	cfg.BaseUrl = "http://localhost:1000/xxx"

	response := `
	[{"workspaceOid":41529001,"subscriptionId":100,"data":{"reason":"manually halted"},"healthCheckShouldFail":false,"timestamp":"2015-02-14T15:32:28Z"}]
	`
	myLoader, passedInRestArgs := test.FakeRestLoader(response)

	var myWriter bytes.Buffer

	args := &showHaltedArgs{
		config: &cfg,
		loader: myLoader,
		writer: &myWriter,
	}

	showHalted(args)

	test.AssertEquals(t, "http://localhost:1000/xxx/info/haltedWorkspaces", passedInRestArgs.Url)

	output := myWriter.String()
	test.AssertContains(t, "Halted Workspaces", output)
	test.AssertContains(t, "41529001", output)
	test.AssertContains(t, "100", output)
	test.AssertContains(t, "manually halted", output)
	test.AssertContains(t, "false", output)
	test.AssertContains(t, "2015-02-14T15:32:28Z", output)
}

