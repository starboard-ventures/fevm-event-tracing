package v1

import (
	"busi/internal/busi/core"
	"busi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// event cron job godoc
// @Description event cron job api, call by dolphin scheduler
// @Tags Inner
// @Accept application/json,json
// @Produce application/json,json
// @Param RequestHeight query core.RequestHeight false "RequestHeight"
// @Success 200 {object} utils.ResponseCode
// @Router /event [post]
func EventHandle(c *gin.Context) {
	app := utils.Gin{C: c}

	var r core.RequestHeight
	if err := c.ShouldBindJSON(&r); err != nil {
		app.HTTPResponse(http.StatusOK, utils.NewResponse(utils.CodeBadRequest, err.Error(), nil))
		return
	}

	resp := core.EventHandle(c.Request.Context(), &r)
	if resp != nil {
		app.HTTPResponse(http.StatusOK, resp.Response)
		return
	}

	app.HTTPResponseOK(nil)
}
