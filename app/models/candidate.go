package models

import (
	"github.com/arthurc0102/gin-vote/db"
	"github.com/gin-gonic/gin/binding"
	"github.com/leebenson/conform"
)

// Candidate table
type Candidate struct {
	ID        int    `json:"id"`
	FirstName string `gorm:"not null;size:50" json:"firstName" conform:"trim"`
	LastName  string `gorm:"not null;size:50" json:"lastName" conform:"trim"`
	Age       uint   `gorm:"not null" json:"age"`
	Politics  string `gorm:"not null;size:500" json:"politics" conform:"trim"`
	Vote      uint   `gorm:"not null;default:0" json:"vote"`
}

// Save model
func (model *Candidate) Save() error {
	if err := model.Validate(); err != nil {
		return err
	}

	return db.Connection.Save(model).Error
}

// Create model
func (model *Candidate) Create() error {
	if err := model.Validate(); err != nil {
		return err
	}

	return db.Connection.Create(model).Error
}

// Delete model
func (model *Candidate) Delete() error {
	if db.Connection.NewRecord(model) {
		return nil
	}

	return db.Connection.Delete(model).Error
}

// Validate model
func (model *Candidate) Validate() error {
	conform.Strings(model)
	return binding.Validator.ValidateStruct(model)
}
