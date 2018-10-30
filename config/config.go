package config

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

// Do all config
func Do(server *gin.Engine) {
	Load()
	Static(server)
}

// Load config from .env file
func Load() {
	_ = godotenv.Load()

	if os.Getenv("GIN_MODE") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
}

// Static file serve setting
func Static(server *gin.Engine) {
	server.Static("/static", "static")
	server.StaticFile("/favicon.ico", "public/favicon.ico")
}
