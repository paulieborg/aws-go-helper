package actions

func (s *Stack) rollback(p ProvisionArgs) (bool) {

	ds := s.Describe(p)

	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}

	return false
}
