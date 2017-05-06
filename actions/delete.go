package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"log"
)

func (c *CF) Delete(stack_name *string) {

	input := cf.DeleteStackInput{StackName: stack_name}
	_, err := c.Service.DeleteStackWithContext(c.Context, &input)

	if err != nil {
		log.Fatal(err)
	} else {
		c.waitDelete(stack_name)
	}
}
