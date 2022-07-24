package freemarketservice

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (service service) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	freeMarket, err := service.usecase.Delete(c.Request.Context(), int64(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, freeMarket)
}
