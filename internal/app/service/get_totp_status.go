package service

import (
	"github.com/go-playground/validator"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/mongo"
)

func GetTotpStatus(req model.GetTotpStatusReq) (model.GetTotpStatusRes, model.ServiceResp) {
	res := model.GetTotpStatusRes{}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return res, model.ServiceError.BadRequestError(err.Error())
	}

	totp, err := mongo.GetInstance().GetTotp(req.AccountId, req.ServiceName)
	if err != nil && err == mongo.ERROR_DATA_NOT_FOUND {
		res.Status = string(config.None)
		return res, model.ServiceError.OK
	}

	res.Status = totp.Status
	return res, model.ServiceError.OK
}
