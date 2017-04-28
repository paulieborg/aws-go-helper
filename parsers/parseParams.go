package parsers

import (
	"encoding/json"
	"bytes"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

// parseParams takes a simple JSON blob of parameters and converts it to a slice
// of CloudFormation parameter structs
func ParseParams(params []byte) (p []*cf.Parameter, err error) {
	var sp map[string]string

	err = json.NewDecoder(bytes.NewReader(params)).Decode(&sp)
	if err != nil {
		return
	}

	for k, v := range sp {
		pkey := k
		pval := v
		p = append(p, &cf.Parameter{
			ParameterKey:   &pkey,
			ParameterValue: &pval,
		})
	}

	return
}