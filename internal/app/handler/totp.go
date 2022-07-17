package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pinyi-lee/core.account.2fa.git/internal/app/service"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
)

// InitTotpHandler init totp
// @Summary init totp
// @Description init totp
// @Tags totp
// @Accept json
// @Produce json
// @Param        body  body      model.InitTotpReq    true  "body"
// @Success      200   {object}  model.InitTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /pviv/2fa/v1/totp/init [post]
func InitTotpHandler(c *gin.Context) {
	req := model.InitTotpReq{}
	c.BindHeader(&req)
	c.BindUri(&req)
	c.ShouldBind(&req)

	res, err := service.InitTotp(req)
	result(c, res, err)
}

// EnableTotpHandler enable totp
// @Summary enable totp
// @Description enable totp
// @Tags totp
// @Accept json
// @Produce json
// @Param        body  body      model.EnableTotpReq    true  "body"
// @Success      200   {object}  model.EnableTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /pviv/2fa/v1/totp/enable [post]
func EnableTotpHandler(c *gin.Context) {
	req := model.EnableTotpReq{}
	c.BindHeader(&req)
	c.BindUri(&req)
	c.ShouldBind(&req)

	res, err := service.EnableTotp(req)
	result(c, res, err)
}

// DisableTotpHandler disable totp
// @Summary disable totp
// @Description disable totp
// @Tags totp
// @Accept json
// @Produce json
// @Param        body  body      model.DisableTotpReq    true  "body"
// @Success      200   {object}  model.DisableTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /pviv/2fa/v1/totp/disable [post]
func DisableTotpHandler(c *gin.Context) {
	req := model.DisableTotpReq{}
	c.BindHeader(&req)
	c.BindUri(&req)
	c.ShouldBind(&req)

	res, err := service.DisableTotp(req)
	result(c, res, err)
}

// VerifyTotpHandler verify totp
// @Summary verify totp
// @Description verify totp
// @Tags totp
// @Accept json
// @Produce json
// @Param        body  body      model.VerifyTotpReq    true  "body"
// @Success      200   {object}  model.VerifyTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /pviv/2fa/v1/totp/verify [post]
func VerifyTotpHandler(c *gin.Context) {
	req := model.VerifyTotpReq{}
	c.BindHeader(&req)
	c.BindUri(&req)
	c.ShouldBind(&req)

	res, err := service.VerifyTotp(req)
	result(c, res, err)
}

// GetTotpStatusHandler get totp status
// @Summary get totp status
// @Description get totp status
// @Tags totp
// @Accept json
// @Produce json
// @Param body body model.GetTotpStatusReq true "body"
// @Success      200   {object}  model.GetTotpStatusRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /pviv/2fa/v1/totp/status [get]
func GetTotpStatusHandler(c *gin.Context) {
	req := model.GetTotpStatusReq{}
	c.BindHeader(&req)
	c.BindUri(&req)
	c.ShouldBind(&req)

	res, err := service.GetTotpStatus(req)
	result(c, res, err)
}
