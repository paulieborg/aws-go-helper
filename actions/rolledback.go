package actions

func rolledback(p_args ProvisionArgs, ) (e bool) {

	ds, _ := Describe(p_args)

	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}

	return false
}
