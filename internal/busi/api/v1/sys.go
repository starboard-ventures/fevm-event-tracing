package v1

import (
	"event-trace/pkg/utils"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Description Healthy examination
// @Tags Sys
// @Accept application/json,json
// @Produce application/json,json
// @Success 200 {string} string "pong"
// @Failure 500 {string} string "error:..."
// @Router /ping [get]
func Ping(c *gin.Context) {
	app := utils.Gin{C: c}
	app.HTTPResponseOK("pong")
}
