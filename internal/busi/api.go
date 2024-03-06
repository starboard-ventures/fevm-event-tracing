package busi

import (
	"context"
	v1 "event-trace/internal/busi/api/v1"
	"event-trace/pkg/models/fevm"
	"event-trace/pkg/utils"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func setDealProposalCreateTrancingConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(v1.LOTUS0, utils.CNF.FevmEvent.Lotus)

		c.Next()
	}
}

func setWfilTrancingConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(v1.LOTUS0, utils.CNF.FevmEvent.Lotus)
		c.Set(v1.WFIL, utils.CNF.FevmEvent.WfilContract)

		c.Next()
	}
}

func setPfilTracingConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(v1.LOTUS0, utils.CNF.FevmEvent.Lotus)
		c.Set(v1.PFIL, utils.CNF.FevmEvent.PfilContract)

		c.Next()
	}
}

func setReplTracingConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(v1.LOTUS0, utils.CNF.FevmEvent.Lotus)
		c.Set(v1.REPL, utils.CNF.FevmEvent.ReplContract)

		c.Next()
	}
}

func setReplAuctionTracingConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(v1.LOTUS0, utils.CNF.FevmEvent.Lotus)
		c.Set(v1.REPLAUCTION, utils.CNF.FevmEvent.ReplAuctionContract)

		c.Next()
	}
}

func RegisterRoutes(r *gin.Engine) {
	// r.Use(utils.Cors())
	r.Use(cors.Default())
	r.GET("/busi/swagger/*any", swagHandler)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
		{
			apiv1.POST("/deal-proposal-create-event-tracing-cron", setDealProposalCreateTrancingConfig(), v1.DealProposalCreateEventCronHandle)
		}

		{
			apiv1.POST("/wfil-event-tracing-cron", setWfilTrancingConfig(), v1.WfilEventCronHandle)
		}

		{
			apiv1.POST("/pfil-event-tracing-cron", setPfilTracingConfig(), v1.PfilEventCronHandle)
			apiv1.POST("/repl-event-tracing-cron", setReplTracingConfig(), v1.ReplEventCronHandle)
			apiv1.POST("/repl-auction-event-tracing-cron", setReplAuctionTracingConfig(), v1.ReplAuctionEventCronHandle)
		}
	}
}

func initconfig(ctx context.Context, cf *utils.TomlConfig) {
	if err := utils.InitConfFile(Flags.Config, cf); err != nil {
		log.Fatalf("Load configuration file err: %v", err)
	}

	if err := utils.InitDBEngine(ctx, cf.FevmEvent.DB, fevm.Tables); err != nil {
		log.Fatalf("Initialize SQL engine err: %v", err)
	}
}

func Start() {
	initconfig(context.Background(), &utils.CNF)

	if Flags.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	RegisterRoutes(r)

	r.Run(utils.CNF.FevmEvent.Addr)
}
