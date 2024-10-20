package app

const packageName = "gateway.app"

type Operations interface {
	Ping() string
}

type App struct {
}

func New() Operations {
	app := &App{}

	return Operations(app)
}

func (a *App) Ping() string {
	return "healthy"
}
