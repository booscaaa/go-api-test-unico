package freemarketservice

import (
	"net/http"
	"strconv"

	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/booscaaa/go-api-test-unico/util"
	"github.com/gin-gonic/gin"
)

// Update goDoc
// @Summary Update free market by id
// @Description Update free market by id
// @Tags free-market
// @Accept  json
// @Produce  json
// @Param freeMarket body dto.FreeMarketRequestBody true "freeMarket"
// @Param id path int true "1"
// @Success 200 {object} domain.FreeMarket
// @Router /free-market/{id} [put]
func (service service) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	freeMarketRequest, err := dto.FromJSONFreeMarketRequestBody(c.Request.Body, &service.validator)
	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	freeMarket, err := service.usecase.Update(c.Request.Context(), int64(id), freeMarketRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, freeMarket)
}
