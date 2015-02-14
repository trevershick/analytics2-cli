package test

import (
	"testing"
	"reflect"
	"strings"
	"encoding/json"
	"github.com/trevershick/analytics2-cli/a2m/rest"
)

/* Test Helpers */
func AssertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", expected, reflect.TypeOf(expected), actual, reflect.TypeOf(actual))
	}
}

func AssertContains(t *testing.T, expected string, output string) {
	if !strings.Contains(output, expected) {
		t.Errorf("Expected '%s' to contain '%s'", output, expected)
	}
}


func FakeRestLoader(responseContent string) (rest.Loader, *rest.RestArgs) {

	passedInRestArgs := rest.RestArgs{}

	myLoader := func(args *rest.RestArgs) (error) {
		// can i do this differently, without this type of assignment?
		// i think i need a pointer to a pointer :)
		passedInRestArgs.Url = args.Url
		passedInRestArgs.Params = args.Params
		err := json.Unmarshal([]byte(responseContent), args.ResponseData)
		if err != nil {
			panic(err)
		}
		return nil
	}

	return myLoader, &passedInRestArgs
}
