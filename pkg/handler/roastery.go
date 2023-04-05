package handler

import (
	"coffeecatalogue4"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createRoastery(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input coffeecatalogue4.Roastery

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Roastery.Create(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateRoastery(c *gin.Context) {

}

func (h *Handler) deleteRoastery(c *gin.Context) {

}

func (h *Handler) getAllRoasteries(c *gin.Context) {

}

func (h *Handler) findRoastery(c *gin.Context) {

}

func (h *Handler) roasteriesCount(c *gin.Context) {

}
