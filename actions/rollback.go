package actions

func (s *StackArgs) rollback() (bool) {

	ds, _ := s.Describe()

	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}

	return false
}
