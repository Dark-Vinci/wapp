package app

import "github.com/dark-vinci/wapp/backend/chats/store"

type Operations interface{}

type App struct {
	channelChat *store.ChannelChat
}

func NewApp() Operations {
	app := &App{}

	return Operations(app)
}
