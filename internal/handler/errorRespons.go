package handler

import (
	"log"
	"test_task/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ErrResp(c *gin.Context, err models.ErrRespons) {
	log.Println(err.Message)
	c.AbortWithStatusJSON(int(err.Code), err)

}
