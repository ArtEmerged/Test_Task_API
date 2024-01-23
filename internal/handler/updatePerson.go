package handler

import (
	"strconv"
	"test_task/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updatePerson(c *gin.Context) {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
		return
	}

	var input models.Person

	err = c.BindJSON(&input)
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

	err = h.services.UpdatePerson(personId, input)
	if err != nil {
		if err == models.ErrNoSuchPerson {
			h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
			return
		}
		h.ErrResp(c, models.ErrRespons{Code: 500, Message: err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})
}
