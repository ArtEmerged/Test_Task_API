package handler

import (
	"strconv"
	"test_task/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getPeople(c *gin.Context) {
	pageNumber, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 400, Message: err.Error()})
		return
	}

	filters := models.Filters{
		Offset: (pageNumber - 1) * pageSize,
		Limit:  pageSize,
		Filters: map[string]interface{}{
			"name":        c.Query("name"),
			"surname":     c.Query("surname"),
			"patronymic":  c.Query("patronymic"),
			"age":         c.Query("age"),
			"gender":      c.Query("gender"),
			"nationalize": c.Query("nationalize"),
		},
	}

	people, err := h.services.GetPeople(filters)
	if err != nil {
		h.ErrResp(c, models.ErrRespons{Code: 500, Message: err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
		"data":   people,
	})

}
