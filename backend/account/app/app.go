package app

type Operations interface{}

type App struct{}

func New() Operations {
	app := &App{}

	return Operations(app)
}
