package parse

import (
	"encoding/json"
	"bytes"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
)

// Params takes a simple JSON blob of parameters and converts it to a slice
// of CloudFormation parameter structs
func Params(params []byte) ([]*cf.Parameter) {
	var sp map[string]string
	var p []*cf.Parameter

	err := json.NewDecoder(bytes.NewReader(params)).Decode(&sp)
	if err != nil {
		panic(err)
	}

	for k, v := range sp {
		p_key := k
		p_val := v
		p = append(p, &cf.Parameter{
			ParameterKey:   &p_key,
			ParameterValue: &p_val,
		})
	}

	return p
}
