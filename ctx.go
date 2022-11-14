package woles

type Parameters map[string]any

type Ctx struct {
	parameters Parameters
}

func (c *Ctx) AddParam(name string, value any) Parameters {
	if c.parameters == nil {
		c.parameters = make(Parameters)
	}

	c.parameters[name] = value

	return c.parameters
}

func (c Ctx) Params(parameterName string) any {
	return c.parameters[parameterName]
}
