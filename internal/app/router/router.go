package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pinyi-lee/core.account.2fa.git/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/pinyi-lee/core.account.2fa.git/internal/app/handler"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
)

var Router *gin.Engine

func SetupRouter() (router *gin.Engine) {
	docs.SwaggerInfo.Version = config.Env.Version

	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.Default()
	router.Use(handler.CORSMiddleware(), handler.ErrorMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/health", handler.HealthHandler)
	router.GET("/version", handler.VersionHandler)

	router.POST("/pviv/2fa/v1/totp/init", handler.InitTotpHandler)
	router.POST("/pviv/2fa/v1/totp/enable", handler.EnableTotpHandler)
	router.POST("/pviv/2fa/v1/totp/disable", handler.DisableTotpHandler)
	router.POST("/pviv/2fa/v1/totp/verify", handler.VerifyTotpHandler)
	router.GET("pviv/2fa/v1/totp/status", handler.GetTotpStatusHandler)

	return
}

func Setup() error {
	Router = SetupRouter()

	return nil
}
