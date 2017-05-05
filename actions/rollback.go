package actions

func (s *StackArgs) rollback() (bool) {

	ds := s.Describe()

	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}

	return false
}
