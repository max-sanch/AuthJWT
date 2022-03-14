package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/max-sanch/AuthJWT"
	"net/http"
)

func (h *Handler) singIn(c *gin.Context) {
	var input auth_jwt.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.services.Authentication.GenerateTokens(input.GUID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, auth_jwt.TokenPair{
		AccessToken: tokens["access_token"],
		RefreshToken: tokens["refresh_token"],
	})
}

func (h *Handler) refresh(c *gin.Context) {
	var input auth_jwt.TokenPair

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.services.Authentication.RefreshTokens(input.AccessToken, input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, auth_jwt.TokenPair{
		AccessToken: tokens["access_token"],
		RefreshToken: tokens["refresh_token"],
	})
}
