package handler

import (
	"coffeecatalogue4/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		roasteries := api.Group("/roasteries")
		{
			roasteries.POST("/", h.createRoastery)
			roasteries.PATCH("/:id", h.updateRoastery)
			roasteries.DELETE("/:id", h.deleteRoastery)
			roasteries.GET("/", h.getAllRoasteries)
			roasteries.GET("/:id", h.findRoastery)

			coffeePacks := api.Group("/coffee_packs")
			{
				coffeePacks.POST("/", h.createCoffeePack)
				coffeePacks.PATCH("/:id", h.updateCoffeePack)
				coffeePacks.DELETE("/:id", h.deleteCoffee)
				coffeePacks.GET("/", h.getAllCoffeePacks)
				coffeePacks.GET("/:id", h.findCoffeePack)
			}
		}
	}

	return router
}
