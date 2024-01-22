package handler

import (
	"strconv"
	"test_task/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deletePerson(c *gin.Context) {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
		return
	}

	err = h.services.DeletePerson(personId)
	if err != nil {
		if err == models.ErrNoSuchPerson {
			h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
			return
		}
		h.ErrResp(c, models.ErrRespons{Code: 500, Message: err.Error()})
		return
	}

	c.JSON(200, map[string]interface{}{
		"status": "ok",
	})
}
