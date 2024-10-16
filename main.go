// package main

// import (
// 	"net/http"

// 	"main/src/api/handler"

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	router := gin.New()
// 	router.Use(gin.Logger())

// 	routerGroup := router.Group("/watchlists")

// 	routerGroup.POST("/", handler.HandleCreateWatchList)

// 	routerGroup.GET("/:id", handler.HandleGetWatchList)

// 	routerGroup.PUT("/:id", handler.HandleUpdateWatchList)

// 	routerGroup.DELETE("/:id", handler.HandleDeleteWatchList)

// 	http.ListenAndServe(":7777", router)

// }

package main

import (
	"main/src/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	router.SetupRoutes(r)

	http.ListenAndServe(":7777", r)
}
