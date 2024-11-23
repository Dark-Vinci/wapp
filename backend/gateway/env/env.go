package env

const packageName = "gateway.env"

type Environment struct {
	FrontEndURL   string
	RedisURL      string
	RedisUsername string
	RedisPassword string
}

func New() *Environment {
	return &Environment{}
}
