package config

import (
	"github.com/arthurc0102/gin-vote/app/actions"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes register route of this project to server
func RegisterRoutes(server *gin.Engine) {
	// Common
	server.GET("", actions.Root)

	// Candidates
	candidates := server.Group("candidates")
	{
		candidates.GET("", actions.GetCandidates)
		candidates.POST("", actions.CreateCandidate)

		candidate := candidates.Group(":id")
		{
			candidate.GET("", actions.GetCandidate)
			candidate.PUT("", actions.UpdateCandidate)
			candidate.DELETE("", actions.DeleteCandidate)
			candidate.PATCH("vote", actions.VoteCandidate)
		}
	}
}
