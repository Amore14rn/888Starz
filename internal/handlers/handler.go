package handlers

import (
	us "github.com/Amore14rn/888Starz/internal/controllers/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router *gin.Engine
}

func NewHandler() *Handler {
	return &Handler{
		router: gin.Default(),
	}
}

func (h *Handler) RegisterRoutes() {
	userGroup := h.router.Group("/user")
	{
		userGroup.POST("/create", us.CreateUser)
		userGroup.GET("/get/:name", us.GetUser)
		userGroup.PATCH("/update", us.UpdateUser)
		userGroup.DELETE("/delete/:name", us.DeleteUser)
		userGroup.POST("/create-order", us.CreateOrder)
		userGroup.POST("/add-to-order", us.AddToOrder)
	}
}
