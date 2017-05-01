package actions

import (
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"log"
)

func (s *StackArgs) Delete() {

	input := cf.DeleteStackInput{StackName: &s.Stack_name}
	_, err := s.Session.DeleteStackWithContext(s.Context, &input)

	if err != nil {
		log.Fatal(err)
	} else {
		s.waitDelete()
	}
}
