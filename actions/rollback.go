package actions

func (c *Context) rollback(p ProvisionArgs) (bool) {

	ds := c.Describe(p)

	if *ds.Stacks[0].StackStatus == "ROLLBACK_COMPLETE" {
		return true
	}

	return false
}
