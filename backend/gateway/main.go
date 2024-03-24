package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	gin.ForceConsoleColor()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":     "tomiwa",
			"response": "200",
		})
	})

	return r
}

func main() {
	_ = os.Setenv("TZ", "Africa/Lagos")

	zlFile, err := os.Create("./zero.log")
	if err != nil {
		panic("cant create file")
	}

	logger := zerolog.New(zlFile).With().Timestamp().Logger()
	appLogger := logger.With().Str("GATEWAY", "api").Logger()

	appLogger.Debug().Msg("something should happen")
	appLogger.Debug().Msg("another log in the logger file")

	r := setupRouter()
	if err := r.Run(":8080"); err != nil {
		appLogger.Err(err).Msg("something went wrong")
	}

	appLogger.Debug().Msg("app axiting")
}
