package actions

func (c *Context) exists(p ProvisionArgs) (bool) {

	ds := c.Describe(p)

	if len(ds.Stacks) > 0 {
		return true
	}

	return false
}
