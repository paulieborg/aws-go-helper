package actions

import (
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"log"
)

func (s *Stack) Delete(stack_name *string) {

	sess := cf.New(session.Must(session.NewSession()))

	input := cf.DeleteStackInput{StackName: stack_name}
	_, err := sess.DeleteStackWithContext(s.Context, &input)

	if err != nil {
		log.Fatal(err)
	} else {
		s.waitDelete(stack_name)
	}
}
