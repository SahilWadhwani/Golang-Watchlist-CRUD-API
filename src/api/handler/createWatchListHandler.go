package handler

import (
	"main/src/api/model"
	"main/src/business"

	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateWatchList(c *gin.Context) {
	var watchlist model.CreateWatchList

	// error handeling
	if err := c.ShouldBindJSON(&watchlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Call the validation function to validate WatchList fields
	if err := model.ValidateCreateWatchList(&watchlist); err != nil {
		// Return validation errors if any
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	// calling the create logic from business,
	business.CreateWatchListService(&watchlist)
	c.JSON(http.StatusCreated, watchlist) // 201 is the status code for Created

}

// Add a new function to get all watchlists for a user
