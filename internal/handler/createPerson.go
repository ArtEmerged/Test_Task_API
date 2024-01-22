package handler

import (
	"test_task/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createPerson(c *gin.Context) {
	var input models.Person

	err := c.BindJSON(&input)
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
		return
	}

	// validation ascii letters
	err = validateData(input)
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
		return
	}

	id, err := h.services.CreatePerson(input)
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
	}

	c.JSON(200, map[string]interface{}{
		"id": id,
	})
}
