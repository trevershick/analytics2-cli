package info

import (
	"bytes"
	"testing"
	"github.com/trevershick/analytics2-cli/a2m/test"
	"github.com/trevershick/analytics2-cli/a2m/config"
)


func Test_showCollectionInformation(t *testing.T) {
	cfg := config.Configuration{}
	cfg.BaseUrl = "http://localhost:1000/xxx"

	response := `
	{
		"collections": [
			{ "name": "coll1", "totalStorageSize": 123456789 },
			{ "name": "coll2", "totalStorageSize": 1234 },
			{ "name": "coll3", "totalStorageSize": 123 }
		]
	}
	`
	myLoader, passedInRestArgs := test.FakeRestLoader(response)

	var myWriter bytes.Buffer

	args := &showCollectionArgs{
		config: &cfg,
		workspaceId: 97,
		loader: myLoader,
		writer: &myWriter,
	}

	showCollectionInformation(args)

	test.AssertEquals(t, "http://localhost:1000/xxx/info/workspace/97", passedInRestArgs.Url)

	output := myWriter.String()
	test.AssertContains(t, "Collections for Workspace 97", output)
	test.AssertContains(t, "coll1", output)
	test.AssertContains(t, "coll2", output)
	test.AssertContains(t, "coll3", output)
	test.AssertContains(t, "123", output)
	test.AssertContains(t, "1.2K", output)
	test.AssertContains(t, "117.7M", output)
}

