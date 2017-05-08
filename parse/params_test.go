package parse

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"testing"
)

func TestParams(t *testing.T) {

	params := []byte(`
		{
			"ParamOne": "1",
			"ParamTwo": "2"
		}`)

	var p []*cf.Parameter = Params(params)

	if !keyInList("ParamOne", p) {
		t.Errorf("expected: %s", "ParamOne")
	}

	if !keyInList("ParamTwo", p) {
		t.Errorf("expected: %s", "ParamTwo")
	}

	if !valueInList("1", p) {
		t.Errorf("expected: %s", "1")
	}

	if !valueInList("2", p) {
		t.Errorf("expected: %s", "2")
	}
}

func TestJsonParserPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("The code did not panic")
		}
	}()

	bad_params := []byte(`
			"I suck and JSON format"
		`)

	Params(bad_params)
}

func keyInList(value string, params []*cf.Parameter) bool {

	for _, v := range params {

		if *v.ParameterKey == value {
			return true
		}
	}
	return false
}

func valueInList(value string, params []*cf.Parameter) bool {

	for _, v := range params {

		if *v.ParameterValue == value {
			return true
		}
	}
	return false
}
