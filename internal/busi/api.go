package busi

import (
	v1 "event-trace/internal/busi/api/v1"
	"event-trace/pkg/models/fevm"
	"event-trace/pkg/utils"
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func setTrancingConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(v1.LOTUS0, utils.CNF.FevmEvent.Lotus)

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
		apiv1.POST("/deal-proposal-create-event-tracing-cron", setTrancingConfig(), v1.DealProposalCreateEventCronHandle)
		apiv1.POST("/deal-proposal-create-event-tracing", setTrancingConfig(), v1.DealProposalCreateEventHandle)
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
