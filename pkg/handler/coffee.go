package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCoffeePack(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateCoffeePack(c *gin.Context) {

}

func (h *Handler) deleteCoffee(c *gin.Context) {

}

func (h *Handler) getAllCoffeePacks(c *gin.Context) {

}

func (h *Handler) findCoffeePack(c *gin.Context) {

}

//create coffeeByRoastery
