package actions

func exists(p_args ProvisionArgs, ) (e bool) {

	ds, _ := Describe(p_args)

	if len(ds.Stacks) > 0 {
		return true
	}

	return false
}
