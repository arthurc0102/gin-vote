package config

import (
	"github.com/arthurc0102/gin-vote/app/validators"
	"github.com/gin-gonic/gin/binding"
	validator "gopkg.in/go-playground/validator.v8"
)

// RegisterValidators register validators
func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notspace", validators.NotSpace)
	}
}
