package router

import (
	"main/src/api/handler"

	"github.com/gin-gonic/gin"
)

func watchListRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/", handler.HandleCreateWatchList)
	routerGroup.GET("/:id", handler.HandleGetWatchList)
	routerGroup.PUT("/:id", handler.HandleUpdateWatchList)
	routerGroup.DELETE("/:id", handler.HandleDeleteWatchList)
}
