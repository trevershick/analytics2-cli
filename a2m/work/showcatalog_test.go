package work

import (
	"bytes"
	"testing"
	"github.com/trevershick/analytics2-cli/a2m/test"
	"github.com/trevershick/analytics2-cli/a2m/config"
)


func Test_showCatalog(t *testing.T) {
	cfg := config.Configuration{}
	cfg.BaseUrl = "http://localhost:1000/xxx"

	response := `
		{"available":[
			{"name":"allAllowedValuesRefreshTask","requiresWorkspaceId":false,"description":"Refresh the allowed values metadata for all workspaces","parameters":[]},
			{"name":"allowedValuesRefreshTask","requiresWorkspaceId":true,"description":"Refresh the allowed values for the provided workspace","parameters":[]},
			{"name":"haltAllWorkspacesInDatabaseTask","requiresWorkspaceId":false,"description":"Halt all workspaces by database partition","parameters":{"databaseToHalt":{"required":false,"type":"int","description":"Database partition 1-30"},"reason":{"required":false,"type":"String","description":"The reason for the halt"}}}
		]}
	`
	myLoader, passedInRestArgs := test.FakeRestLoader(response)

	var myWriter bytes.Buffer

	args := &showCatalogArgs{
		config: &cfg,
		loader: myLoader,
		writer: &myWriter,
	}

	showCatalog(args)

	test.AssertEquals(t, "http://localhost:1000/xxx/management/work/catalog", passedInRestArgs.Url)

	output := myWriter.String()
	test.AssertContains(t, "Available Tasks in the Work Catalog", output)
	test.AssertContains(t, "allAllowedValuesRefreshTask", output)
	test.AssertContains(t, "allowedValuesRefreshTask", output)
	test.AssertContains(t, "haltAllWorkspacesInDatabaseTask", output)
	test.AssertContains(t, "Halt all workspaces by database partition", output)
	test.AssertContains(t, "Refresh the allowed values for the provided workspace", output)
	test.AssertContains(t, "Refresh the allowed values metadata for all workspaces", output)
}

