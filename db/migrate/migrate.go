package migrate

import (
	"github.com/arthurc0102/gin-vote/app/models"
	"github.com/arthurc0102/gin-vote/db"
)

// Migrate to db
func Migrate() {
	db.Connection.AutoMigrate(
		&models.Candidate{},
	)
}
