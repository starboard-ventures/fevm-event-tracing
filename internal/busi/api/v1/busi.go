package v1

import (
	"event-trace/internal/busi/core"
	"event-trace/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// deal-proposal-create's event cron job godoc
// @Description deal-proposal-create's event cron job api, call by dolphin scheduler
// @Tags Inner|Cron
// @Accept application/json,json
// @Produce application/json,json
// @Success 200 {object} utils.ResponseCode
// @Router /deal-proposal-create-event-tracing-cron [post]
func DealProposalCreateEventCronHandle(c *gin.Context) {
	app := utils.Gin{C: c}

	lotus0, _ := c.Get(LOTUS0)
	lotus0Cfg, _ := lotus0.(string)

	resp := core.DealProposalCreateEventCronHandle(c.Request.Context(), lotus0Cfg)
	if resp != nil {
		app.HTTPResponse(http.StatusOK, resp.Response)
		return
	}

	app.HTTPResponseOK(nil)
}

// deal-proposal-create's event manual job godoc
// @Description deal-proposal-create's event manual job api
// @Tags Inner|Manual
// @Accept application/json,json
// @Produce application/json,json
// @Param RequestHeight query core.RequestHeight false "RequestHeight"
// @Success 200 {object} utils.ResponseCode
// @Router /deal-proposal-create-event-tracing [post]
func DealProposalCreateEventHandle(c *gin.Context) {
	app := utils.Gin{C: c}

	var r core.RequestHeight
	if err := c.ShouldBindQuery(&r); err != nil {
		app.HTTPResponse(http.StatusOK, utils.NewResponse(utils.CodeBadRequest, err.Error(), nil))
		return
	}

	if err := r.Validate(); err != nil {
		app.HTTPResponse(http.StatusBadRequest, utils.NewResponse(utils.CodeBadRequest, err.Error(), nil))
		return
	}

	lotus0, _ := c.Get(LOTUS0)
	r.Lotus0, _ = lotus0.(string)

	resp := core.DealProposalCreateEventHandle(c.Request.Context(), &r)
	if resp != nil {
		app.HTTPResponse(http.StatusOK, resp.Response)
		return
	}

	app.HTTPResponseOK(nil)
}

// wfil's event cron job godoc
// @Description wfil's event cron job api, call by dolphin scheduler
// @Tags Inner|Cron
// @Accept application/json,json
// @Produce application/json,json
// @Success 200 {object} utils.ResponseCode
// @Router /wfil-event-tracing-cron [post]
func WfilEventCronHandle(c *gin.Context) {
	app := utils.Gin{C: c}

	lotus0, _ := c.Get(LOTUS0)
	lotus0Cfg, _ := lotus0.(string)

	wfil, _ := c.Get(WFIL)
	wfilContract, _ := wfil.(string)

	resp := core.WfilEventCronHandle(c.Request.Context(), lotus0Cfg, strings.ToLower(wfilContract))
	if resp != nil {
		app.HTTPResponse(http.StatusOK, resp.Response)
		return
	}

	app.HTTPResponseOK(nil)
}
