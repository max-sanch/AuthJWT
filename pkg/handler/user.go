package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/max-sanch/AuthJWT"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var input auth_jwt.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	guid, err := h.services.User.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"guid": guid,
	})
}

func (h *Handler) deleteUser(c *gin.Context) {
	var input auth_jwt.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	guid, err := h.services.User.DeleteUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"guid": guid,
	})
}
