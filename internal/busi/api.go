package busi

import (
	v1 "busi/internal/busi/api/v1"
	"busi/pkg/models/fevm"
	"busi/pkg/utils"
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func RegisterRoutes(r *gin.Engine) {
	// r.Use(utils.Cors())
	r.Use(cors.Default())
	r.GET("/busi/swagger/*any", swagHandler)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/ping", v1.Ping)
		r.POST("/event", v1.EventHandle)
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
