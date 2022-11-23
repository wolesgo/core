package router

type ArgumentConfig struct {
	required  bool
	optional  bool
	mandatory bool
}

type Argument struct {
	description string
}

func NewArgument(config ArgumentConfig, name string, description string) Argument {
	return Argument{
		description: description,
	}
}
