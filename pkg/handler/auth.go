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
}

func (h *Handler) refresh(c *gin.Context) {

}
