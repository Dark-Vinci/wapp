package app

const packageName = "gateway.app"

type Operations interface {
	Ping() string
	//CreateAccount() string
	//Login() string
	//Verify phone number
	//2FA
	//VERIFY 2FA

	// SECRET

	// GALLERY

	//
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
