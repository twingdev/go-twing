package plugin

type Environment struct {
	Repo   string
	Config interface{}
}

type Plugin interface {
	Name() string
	Version() string
	Signed() bool
	Init(env *Environment) error
}
