package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	people := router.Group("/people")
	{
		people.GET("/", h.getPeople)
		people.GET("/:id", h.getPerson)
		people.POST("/", h.createPerson) //+
		people.PUT("/:id", h.updatePerson) //+
		people.DELETE("/:id", h.deletePerson) //+
	}

	return router
}
