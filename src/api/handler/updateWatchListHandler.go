package handler

import (
	"main/src/api/model"
	"main/src/business"

	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleUpdateWatchList(c *gin.Context) {
	id := c.Param("id")
	var watchlist model.UpdateWatchList

	// error handeling
	if err := c.ShouldBindJSON(&watchlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Call the validation function to validate WatchList fields
	if err := model.ValidateUpdateWatchList(&watchlist); err != nil {
		// Return validation errors if any
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	watchlistUpdated, err := business.UpdateWatchListService(id, &watchlist) // calling the update logic from business, UpdateWatchListService(id, &watchlist)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, watchlistUpdated)

}
