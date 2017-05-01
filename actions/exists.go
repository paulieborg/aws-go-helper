package actions

func (s *StackArgs) exists() (bool) {

	ds := s.Describe()

	if len(ds.Stacks) > 0 {
		return true
	}

	return false
}
