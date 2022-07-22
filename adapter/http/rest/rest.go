package rest

import (
	"github.com/gin-gonic/gin"
)

func InitRest(httpPort string) {
	router := gin.Default()

	router.Run(":" + httpPort)
}
