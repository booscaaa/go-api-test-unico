package freemarketservice

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Delete goDoc
// @Summary Delete free market by id
// @Description Delete free market by id
// @Tags free-market
// @Accept  json
// @Produce  json
// @Param id path int true "1"
// @Success 200 {object} domain.FreeMarket
// @Router /free-market/{id} [delete]
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
