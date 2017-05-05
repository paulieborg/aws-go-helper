package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"log"
)

func (c *Context) Delete(stack_name *string) {

	sess := cf.New(session.Must(session.NewSession()))

	input := cf.DeleteStackInput{StackName: stack_name}
	_, err := sess.DeleteStackWithContext(c.Context, &input)

	if err != nil {
		log.Fatal(err)
	} else {
		c.waitDelete(stack_name)
	}
}
