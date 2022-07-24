package freemarketservice

import (
	"net/http"

	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/gin-gonic/gin"
)

// Fetch goDoc
// @Summary Fetch free markets with server pagination
// @Description Fetch free markets with server pagination
// @Tags free-market
// @Accept  json
// @Produce  json
// @Param sort query string false "column_name1,column_name2"
// @Param descending query string false "true,false"
// @Param page query integer true "1"
// @Param itemsPerPage query integer true "10"
// @Param search query string false "string word, can be empty"
// @Success 200 {object} dto.PaginationResponseBody
// @Router /free-market [get]
func (service service) Fetch(c *gin.Context) {
	paginationRequest := dto.FromValuePaginationRequestParams(c.Request)

	freeMarket, err := service.usecase.Fetch(c.Request.Context(), paginationRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, freeMarket)
}
