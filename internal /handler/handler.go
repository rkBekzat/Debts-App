package handler

import (
	"debts/internal /service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{serv: serv}
}

func (h *Handler) InitRoute(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
}
