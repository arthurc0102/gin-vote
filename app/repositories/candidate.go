package repositories

import (
	"github.com/arthurc0102/gin-vote/app/models"
	"github.com/arthurc0102/gin-vote/db"
)

// GetCandidates return all candidate
func GetCandidates() (candidates []models.Candidate) {
	db.Connection.Find(&candidates)
	return
}

// GetCandidateByID return candidate filter by id
func GetCandidateByID(id interface{}) (candidate models.Candidate, exists bool) {
	db.Connection.First(&candidate, id)
	exists = !db.Connection.NewRecord(candidate)
	return
}
