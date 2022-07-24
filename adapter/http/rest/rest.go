package rest

import (
	"github.com/booscaaa/go-api-test-unico/di"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRest(httpPort string, database *sqlx.DB) {
	freeMarketService := di.ConfigFreeMarketDi(database)

	router := gin.Default()

	router.POST("/free-market", freeMarketService.Create)
	router.GET("/free-market", freeMarketService.Fetch)
	router.GET("/free-market/:id", freeMarketService.GetByID)
	router.PUT("/free-market/:id", freeMarketService.Update)
	router.DELETE("/free-market/:id", freeMarketService.Delete)

	router.Run(":" + httpPort)
}
