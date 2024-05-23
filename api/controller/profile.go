package controller

import (
	"net/http"
	"send2kobo/domain"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

// @Summary Fetch user profile
// @Tags User
// @Description  查询用户信息
// @Security token
// @Accept  json
// @Produce  json
// @Success 200 {object} domain.Profile
// @Failure 200 {object} domain.ErrorResponse
// @Router /api/v1/profile [get]
func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
