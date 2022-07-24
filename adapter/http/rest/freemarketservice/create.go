package freemarketservice

import (
	"net/http"

	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/booscaaa/go-api-test-unico/util"
	"github.com/gin-gonic/gin"
)

func (service service) Create(c *gin.Context) {
	freeMarketRequest, err := dto.FromJSONFreeMarketRequestBody(c.Request.Body, &service.validator)
	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	freeMarket, err := service.usecase.Create(c.Request.Context(), freeMarketRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, freeMarket)
}
