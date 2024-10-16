package handler

import (
	"main/src/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleDeleteWatchList(c *gin.Context) {
	id := c.Param("id")

	if err := business.DeleteWatchListService(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})

}
