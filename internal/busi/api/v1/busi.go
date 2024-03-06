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

// pfil's event cron job godoc
// @Description pfil's event cron job api, call by dolphin scheduler
// @Tags Inner|Cron
// @Accept application/json,json
// @Produce application/json,json
// @Success 200 {object} utils.ResponseCode
// @Router /pfil-event-tracing-cron [post]
func PfilEventCronHandle(c *gin.Context) {
	app := utils.Gin{C: c}

	lotus0, _ := c.Get(LOTUS0)
	lotus0Cfg, _ := lotus0.(string)

	pfil, _ := c.Get(PFIL)
	pfilContract, _ := pfil.(string)

	resp := core.PfilEventCronHandle(c.Request.Context(), lotus0Cfg, strings.ToLower(pfilContract))
	if resp != nil {
		app.HTTPResponse(http.StatusOK, resp.Response)
		return
	}

	app.HTTPResponseOK(nil)
}

// repl's event cron job godoc
// @Description repl's event cron job api, call by dolphin scheduler
// @Tags Inner|Cron
// @Accept application/json,json
// @Produce application/json,json
// @Success 200 {object} utils.ResponseCode
// @Router /repl-event-tracing-cron [post]
func ReplEventCronHandle(c *gin.Context) {
	app := utils.Gin{C: c}

	lotus0, _ := c.Get(LOTUS0)
	lotus0Cfg, _ := lotus0.(string)

	repl, _ := c.Get(REPL)
	replContract, _ := repl.(string)

	resp := core.ReplEventCronHandle(c.Request.Context(), lotus0Cfg, strings.ToLower(replContract))
	if resp != nil {
		app.HTTPResponse(http.StatusOK, resp.Response)
		return
	}

	app.HTTPResponseOK(nil)
}
