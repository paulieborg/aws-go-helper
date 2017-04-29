package actions

func (p_args *ProvisionArgs) rollback() (e bool) {

	ds, _ := p_args.Describe()

	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}

	return false
}
