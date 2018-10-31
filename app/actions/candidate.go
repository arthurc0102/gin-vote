package actions

import (
	"net/http"

	"github.com/arthurc0102/gin-vote/app/models"
	"github.com/arthurc0102/gin-vote/app/repositories"
	"github.com/arthurc0102/gin-vote/app/serializers"
	"github.com/arthurc0102/gin-vote/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// GetCandidates list all candidates
func GetCandidates(c *gin.Context) {
	candidates := repositories.GetCandidates()
	c.JSON(http.StatusOK, candidates)
}

// CreateCandidate create candidate
func CreateCandidate(c *gin.Context) {
	var serializer serializers.Candidate

	if err := c.ShouldBind(&serializer); err != nil {
		errors := utils.HandleError(serializer, err)
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	candidate := models.Candidate{}
	copier.Copy(&candidate, &serializer)

	if err := candidate.Save(); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleError(candidate, err))
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

// GetCandidate return candidate of id
func GetCandidate(c *gin.Context) {
	candidate, exists := repositories.GetCandidateByID(c.Param("id"))

	if !exists {
		utils.NotFound(c)
		return
	}

	c.JSON(http.StatusOK, candidate)
}

// UpdateCandidate update candidate
func UpdateCandidate(c *gin.Context) {
	candidate, exists := repositories.GetCandidateByID(c.Param("id"))

	if !exists {
		utils.NotFound(c)
		return
	}

	serializer := serializers.Candidate{}

	if err := c.ShouldBind(&serializer); err != nil {
		errors := utils.HandleError(serializer, err)
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	copier.Copy(&candidate, &serializer)

	if err := candidate.Save(); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleError(candidate, err))
		return
	}

	c.JSON(http.StatusCreated, candidate)
}

// DeleteCandidate delete candidate
func DeleteCandidate(c *gin.Context) {
	candidate, exists := repositories.GetCandidateByID(c.Param("id"))

	if !exists {
		utils.NotFound(c)
		return
	}

	if err := candidate.Delete(); err != nil {
		c.JSON(http.StatusBadRequest, utils.HandleError(candidate, err))
		return
	}

	c.JSON(http.StatusNoContent, "")
}

// VoteCandidate vote candidate
func VoteCandidate(c *gin.Context) {
	candidate, exists := repositories.GetCandidateByID(c.Param("id"))

	if !exists {
		utils.NotFound(c)
		return
	}

	candidate.Vote++
	candidate.Save()

	c.JSON(http.StatusOK, candidate)
}
