package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"log"
)

func Delete(stack *StackArgs) error {

	input := cf.DeleteStackInput{StackName: &stack.Stack_name}
	_, err := stack.Session.DeleteStackWithContext(stack.Context, &input)

	if err != nil {
		log.Fatal(err)
	} else {
		err = waitDelete(stack)
	}
	return err
}
