package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	watchListGroup := r.Group("/watchlists")
	watchListRoutes(watchListGroup)
}
