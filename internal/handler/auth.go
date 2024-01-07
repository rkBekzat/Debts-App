package handler

import (
	"debts/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.serv.Auth.SignUp(user)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"id": id,
	})
}

func (h *Handler) SignIn(ctx *gin.Context) {

}
