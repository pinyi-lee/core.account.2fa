package service

import (
	"github.com/go-playground/validator"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/mongo"
)

func EnableTotp(req model.EnableTotpReq) (model.EnableTotpRes, model.ServiceResp) {
	res := model.EnableTotpRes{}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return res, model.ServiceError.BadRequestError(err.Error())
	}

	totp, err := mongo.GetInstance().GetTotp(req.AccountId, req.ServiceName)
	if err != nil && err == mongo.ERROR_DATA_NOT_FOUND {
		return res, model.ServiceError.BadRequestError(err.Error())
	}
	if err != nil {
		return res, model.ServiceError.InternalServiceError(err.Error())
	}

	if totp.Status == string(config.Created) {
		return res, model.ServiceError.StatusConflictError("totp already created")
	}

	valid := Verify(req.Passcode, totp.Secret)
	if valid == false {
		return res, model.ServiceError.BadRequestError("passcode verify fail")
	}

	err = mongo.GetInstance().UpdateTotp(req.AccountId, req.ServiceName, string(config.Created))
	if err != nil {
		return res, model.ServiceError.InternalServiceError(err.Error())
	}

	return res, model.ServiceError.OK
}
