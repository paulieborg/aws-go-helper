package actions

func (p_args *ProvisionArgs) exists() (e bool) {

	ds, _ := p_args.Describe()

	if len(ds.Stacks) > 0 {
		return true
	}

	return false
}
