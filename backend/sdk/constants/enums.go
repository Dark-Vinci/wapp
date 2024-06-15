package constants

type AppEnvironment int

const (
	Testing = iota
	Development
	Production
)

var env = map[AppEnvironment]string {
	Testing: "tesing",
	Development: "development",
	Production: "production",
}

func (a AppEnvironment) String() string {
	return env[a]
}

func FromStr(val string) AppEnvironment {
	for k, v := range env {
		if v == val {
			return k
		}
	}
	
	return 0
}