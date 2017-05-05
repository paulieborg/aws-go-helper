package actions

func (s *Stack) exists(p ProvisionArgs) (bool) {

	ds := s.Describe(p)

	if len(ds.Stacks) > 0 {
		return true
	}

	return false
}
