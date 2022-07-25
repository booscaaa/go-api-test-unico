package freemarketservice

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetByID goDoc
// @Summary Get free market by id
// @Description Get free market by id
// @Tags free-market
// @Accept  json
// @Produce  json
// @Param id path int true "1"
// @Success 200 {object} domain.FreeMarket
// @Router /free-market/{id} [get]
func (service service) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		service.logger.Error("Param ID with problem", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	freeMarket, err := service.usecase.GetByID(c.Request.Context(), int64(id))

	if err != nil {
		service.logger.Error("Erro on get free market by id", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, freeMarket)
}
