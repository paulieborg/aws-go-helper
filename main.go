package main

import (
	"context"
	"flag"
	"fmt"

	"io/ioutil"

	"github.com/paulieborg/aws-go-helper/actions"
	"github.com/paulieborg/aws-go-helper/parse"
	"log"
)

var (
	ctx = context.Background()
)

var (
	action   = flag.String("a", "provision", "create or delete")
	name     = flag.String("n", "", "Stack name.")
	template = flag.String("t", "templates/test-template.yml", "Template file path.")
	params   = flag.String("p", "templates/test-params.json", "Parameters file path.")
	bucket   = flag.String("b", "", "Bucket containing template.")
	timeout  = flag.Int64("x", 5, "Timeout in minutes.")
)

func main() {
	flag.Parse()

	c := actions.Context{
		Context: ctx,
	}

	args := actions.ProvisionArgs{
		Stack_name: *name,
		Parameters: parse.Params(readFile(*params)),
		Template:   readFile(*template),
		BucketName: *bucket,
		Timeout:    *timeout,
	}

	switch *action {
	case "provision":
		status := c.Provision(args)
		fmt.Printf("Stack - %s\n", status)
	case "delete":
		c.Delete(name)
	default:
		fmt.Printf("Unknown action '%s'\n", *action)
	}
}

func readFile(f string) ([]byte) {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
