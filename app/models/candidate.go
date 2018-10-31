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
func (model *Candidate) Save(check ...bool) error {
	if len(check) == 0 || !check[0] {
		model.Conform()
		if err := model.Validate(); err != nil {
			return err
		}
	}

	if db.Connection.NewRecord(model) {
		return db.Connection.Create(model).Error
	}

	return db.Connection.Save(model).Error
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
	return binding.Validator.ValidateStruct(model)
}

// Conform model's string
func (model *Candidate) Conform() {
	conform.Strings(model)
}
