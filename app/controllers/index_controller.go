package controllers

import (
	"net/http"

	"github.com/KCKT0112/LiteBlog/app/services"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	indexService services.IndexService
}

func NewIndexController(service services.IndexService) *IndexController {
	return &IndexController{
		indexService: service,
	}
}

// GetIndex
// @Summary Index
// @Description Index
// @Tags    Index
// @Accept  text/html
// @Produce  text/html
// @Success 200
// @Failure 400
// @Router / [get]
func (uc *IndexController) GetIndex(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("it works!"))
}
