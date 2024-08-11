package env

const packageName = "gateway.env"

type Environment struct {
	FrontEndURL string
}

func New() *Environment {
	return &Environment{}
}
