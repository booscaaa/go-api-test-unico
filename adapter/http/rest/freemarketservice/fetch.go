package freemarketservice

import (
	"net/http"

	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/gin-gonic/gin"
)

func (service service) Fetch(c *gin.Context) {
	paginationRequest := dto.FromValuePaginationRequestParams(c.Request)

	freeMarket, err := service.usecase.Fetch(c.Request.Context(), paginationRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, freeMarket)
}
