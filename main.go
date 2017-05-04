package main

import (
	"context"
	"flag"
	"fmt"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/paulieborg/aws-go-helper/actions"
	"github.com/paulieborg/aws-go-helper/parsers"
	"log"
)

var (
	ctx = context.Background()
	svc = cf.New(session.Must(session.NewSession()))
)

var (
	action   = flag.String("a", "create", "create or delete")
	name     = flag.String("n", "", "Stack name.")
	template = flag.String("t", "network/test-template.yml", "Template file path.")
	params   = flag.String("p", "network/test-params.json", "Parameters file path.")
	bucket   = flag.String("b", "", "Bucket containing template.")
	timeout  = flag.Int64("x", 5, "Timeout in minutes.")
)

func main() {
	flag.Parse()

	stack := actions.StackArgs{
		Context:    ctx,
		Session:    svc,
		Parameters: parsers.ParseParams(readFile(*params)),
		Stack_name: *name,
		Template:   readFile(*template),
		Bucket:     *bucket,
		Timeout:    *timeout, }

	switch *action {
	case "provision":
		d := actions.Provision(&stack)
		fmt.Printf("Stack - %s\n", *d)

	case "delete":
		actions.Delete(&stack)
		fmt.Println("Stack - Deleted")
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
