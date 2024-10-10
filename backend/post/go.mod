module github.com/dark-vinci/wapp/backend/post

replace github.com/dark-vinci/wapp/backend/sdk => ../sdk

go 1.22.1

require (
	github.com/dark-vinci/wapp/backend/sdk v0.0.0
	github.com/pressly/goose/v3 v3.21.1
	github.com/rs/zerolog v1.33.0
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
	gorm.io/plugin/dbresolver v1.5.2
)
