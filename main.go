package main

import (
	"time"

	"github.com/arthurc0102/gin-vote/config"
	"github.com/arthurc0102/gin-vote/db"
	"github.com/arthurc0102/gin-vote/db/migrate"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	config.Load()
	config.Static(server)
	config.RegisterRoutes(server)
	config.RegisterValidators()

	db.Connect()
	defer db.Close()

	migrate.Migrate()

	server.Run()
}
