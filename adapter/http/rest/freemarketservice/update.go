package freemarketservice

import (
	"net/http"
	"strconv"

	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/booscaaa/go-api-test-unico/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
		service.logger.Error("Param ID with problem", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	freeMarketRequest, err := dto.FromJSONFreeMarketRequestBody(c.Request.Body, &service.validator)
	if err != nil {
		if _, ok := err.(util.RequestError); ok {
			service.logger.Error("Erro on json struct fields", zap.Error(err.(util.RequestError)))
			c.JSON(http.StatusUnprocessableEntity, err.(util.RequestError))
			return
		}

		service.logger.Error("Erro on convert json to DTO", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	freeMarket, err := service.usecase.Update(c.Request.Context(), int64(id), freeMarketRequest)

	if err != nil {
		service.logger.Error("Erro on update a new free market", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, freeMarket)
}
