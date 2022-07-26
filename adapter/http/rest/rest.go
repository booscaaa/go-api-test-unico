package rest

import (
	"github.com/booscaaa/go-api-test-unico/di"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	docs "github.com/booscaaa/go-api-test-unico/adapter/http/rest/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRest(httpPort string, database *sqlx.DB, logger *zap.Logger) {
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Host = viper.GetString("server.http.host")
	docs.SwaggerInfo.Schemes = []string{viper.GetString("server.http.scheme")}
	docs.SwaggerInfo.Title = "UNICO API TEST"
	freeMarketService := di.ConfigFreeMarketDi(database, logger)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Use(Logger(logger))

	router.POST("/free-market", freeMarketService.Create)
	router.GET("/free-market", freeMarketService.Fetch)
	router.GET("/free-market/:id", freeMarketService.GetByID)
	router.PUT("/free-market/:id", freeMarketService.Update)
	router.DELETE("/free-market/:id", freeMarketService.Delete)

	router.Run(":" + httpPort)
}
