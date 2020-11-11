package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-swag-demo-api/pkg/models"
)

type PingController struct{}

// Status godoc
// @Summary Returns a 200
// @Accept json
// @Produce json
// @Success 200 {object} models.Status
// @Failure 400
// @Router /v1/ping [get]
func (h PingController) Status(c *gin.Context) {
	status := models.Status{Message: "Pong!"}
	c.JSON(http.StatusOK, status)
}
