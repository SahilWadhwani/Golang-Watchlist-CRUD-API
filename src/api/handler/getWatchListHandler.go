package handler

import (
	"main/src/business"

	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetWatchList(c *gin.Context) {
	id := c.Param("id")
	watchlist, err := business.GetWatchListService(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{ // status code is 404
			"error": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, watchlist)
}
