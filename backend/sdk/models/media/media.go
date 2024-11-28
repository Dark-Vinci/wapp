package media

type Media string

var (
	ForProfile Media = "profile"
	ForChannel Media = "channel"
	ForChat    Media = "chat"
	ForPost    Media = "post"
)
