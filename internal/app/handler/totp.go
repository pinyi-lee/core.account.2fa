package handler

import (
	"github.com/gin-gonic/gin"
)

// InitTotpHandler init totp
// @Summary init totp
// @Description init totp
// @Tags totp
// @Accept json
// @Produce json
// @param account-id header string true "account-id"
// @Param        body  body      model.InitTotpReq    true  "body"
// @Success      200   {object}  model.InitTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /api/2fa/v1/totp/init [post]
func InitTotpHandler(c *gin.Context) {

}

// EnableTotpHandler enable totp
// @Summary enable totp
// @Description enable totp
// @Tags totp
// @Accept json
// @Produce json
// @param account-id header string true "account-id"
// @Param        body  body      model.EnableTotpReq    true  "body"
// @Success      200   {object}  model.EnableTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /api/2fa/v1/totp/enable [post]
func EnableTotpHandler(c *gin.Context) {

}

// DisableTotpHandler disable totp
// @Summary disable totp
// @Description disable totp
// @Tags totp
// @Accept json
// @Produce json
// @param account-id header string true "account-id"
// @Param        body  body      model.DisableTotpReq    true  "body"
// @Success      200   {object}  model.DisableTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /api/2fa/v1/totp/disable [post]
func DisableTotpHandler(c *gin.Context) {

}

// VerifyTotpHandler verify totp
// @Summary verify totp
// @Description verify totp
// @Tags totp
// @Accept json
// @Produce json
// @param account-id header string true "account-id"
// @Param        body  body      model.VerifyTotpReq    true  "body"
// @Success      200   {object}  model.VerifyTotpRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /api/2fa/v1/totp/verify [post]
func VerifyTotpHandler(c *gin.Context) {

}

// GetTotpStatusHandler get totp status
// @Summary get totp status
// @Description get totp status
// @Tags totp
// @Accept json
// @Produce json
// @param account-id header string true "account-id"
// @Param body body model.GetTotpStatusReq true "body"
// @Success      200   {object}  model.GetTotpStatusRes{}  "ok"
// @Failure  204  {object}  model.ServiceErrMsg  "no content"
// @Failure  400  {object}  model.ServiceErrMsg  "bad request"
// @Failure  424  {object}  model.ServiceErrMsg  "failed dependency"
// @Failure  500  {object}  model.ServiceErrMsg  "internal server error"
// @Router /api/2fa/v1/totp/status [get]
func GetTotpStatusHandler(c *gin.Context) {

}
