package parsers

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"testing"
)

func TestParseParams(t *testing.T) {

	params := []byte(`
		{
			"ParamOne": "1",
			"ParamTwo": "2"
		}`)

	var p []*cf.Parameter = ParseParams(params)

	if ! isParameterKeyInList("ParamOne", p) {
		t.Errorf("expected: %s", "ParamOne")
	}

	if ! isParameterKeyInList("ParamTwo", p) {
		t.Errorf("expected: %s", "ParamTwo")
	}

	if ! isParameterValueInList("1", p) {
		t.Errorf("expected: %s", "1")
	}

	if ! isParameterValueInList("2", p) {
		t.Errorf("expected: %s", "2")
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	bad_params := []byte(`
			"I suck and JSON format"
		`)

	ParseParams(bad_params)
}

func isParameterKeyInList(value string, params []*cf.Parameter) bool {

	for _, v := range params {

		if *v.ParameterKey == value {
			return true
		}
	}
	return false
}

func isParameterValueInList(value string, params []*cf.Parameter) bool {

	for _, v := range params {

		if *v.ParameterValue == value {
			return true
		}
	}
	return false
}
