package handler

import (
	"net/http"
	"strconv"
	"test_task/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getPerson(c *gin.Context) {
	personId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
		return
	}

	person, err := h.services.GetPersonById(personId)
	if err != nil {
		if err == models.ErrNoSuchPerson {
			h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
			return
		}
		h.ErrResp(c, models.ErrRespons{Code: 500, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}
