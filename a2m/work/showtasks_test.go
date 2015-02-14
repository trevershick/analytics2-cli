package work

import (
	"bytes"
	"testing"
	"github.com/trevershick/analytics2-cli/a2m/test"
	"github.com/trevershick/analytics2-cli/a2m/config"
)


func Test_showTasks(t *testing.T) {
	cfg := config.Configuration{}
	cfg.BaseUrl = "http://localhost:1000/xxx"

	response := `{
		"tasks":[
			{"cancellable":false,"startDate":"2015-02-14T19:57:00Z","starting":false,"done":true,"highPriority":true,"totalSteps":-1,"cancelled":false,"etaInSeconds":-1,"parameters":{},"started":false,"taskKey":"revisionQueueTask","taskName":"revisionQueueTask","cancellationRequested":false,"currentMessage":null,"status":"COMPLETED","completed":true,"failed":false,"failureReason":null,"endDate":"2015-02-14T19:57:00Z","percentDone":-1,"currentStep":-1,"requeueRequested":false,"workspaceId":null,"id":"revisionQueueTask"}
		]
		,"pageSize":25
		,"page":0
		,"total":1
	}
	`
	myLoader, passedInRestArgs := test.FakeRestLoader(response)

	var myWriter bytes.Buffer

	args := &showTasksArgs{
		config: &cfg,
		loader: myLoader,
		writer: &myWriter,
	}

	showTasks(args)

	test.AssertEquals(t, "http://localhost:1000/xxx/management/work/tasks", passedInRestArgs.Url)

	output := myWriter.String()
	test.AssertContains(t, "Active Tasks", output)
	test.AssertContains(t, "2015-02-14T19:57:00Z", output)
	test.AssertContains(t, "2015-02-14T19:57:00Z", output)
	test.AssertContains(t, "COMPLETED", output)
	test.AssertContains(t, "revisionQueueTask", output)
}

func Test_showRecentTasksAddsFlagToParams(t *testing.T) {
	cfg := config.Configuration{}
	cfg.BaseUrl = "http://localhost:1000/xxx"

	response := `{
		"tasks":[
			{"cancellable":false,"startDate":"2015-02-14T19:57:00Z","starting":false,"done":true,"highPriority":true,"totalSteps":-1,"cancelled":false,"etaInSeconds":-1,"parameters":{},"started":false,"taskKey":"revisionQueueTask","taskName":"revisionQueueTask","cancellationRequested":false,"currentMessage":null,"status":"COMPLETED","completed":true,"failed":false,"failureReason":null,"endDate":"2015-02-14T19:57:00Z","percentDone":-1,"currentStep":-1,"requeueRequested":false,"workspaceId":null,"id":"revisionQueueTask"}
		]
		,"pageSize":25
		,"page":0
		,"total":1
	}`
	myLoader, passedInRestArgs := test.FakeRestLoader(response)

	var myWriter bytes.Buffer

	args := &showTasksArgs{
		recent: true,
		config: &cfg,
		loader: myLoader,
		writer: &myWriter,
	}

	showTasks(args)

	test.AssertEquals(t, "http://localhost:1000/xxx/management/work/tasks", passedInRestArgs.Url)
	test.AssertEquals(t, "false", passedInRestArgs.Params.Get("active"))
}

