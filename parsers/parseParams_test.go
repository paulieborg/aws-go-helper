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

	var p []*cf.Parameter
	p = ParseParams(params)

	if *p[0].ParameterKey != "ParamOne" {
		t.Errorf("expected %s, actual %s", "ParamOne", *p[0].ParameterKey)
	}

	if *p[0].ParameterValue != "1" {
		t.Errorf("expected %s, actual %s", "1", *p[0].ParameterValue)
	}

	if *p[1].ParameterKey != "ParamTwo" {
		t.Errorf("expected %s, actual %s", "ParamTwo", *p[1].ParameterKey)
	}

	if *p[1].ParameterValue != "2" {
		t.Errorf("expected %s, actual %s", "2", *p[1].ParameterValue)
	}

}
