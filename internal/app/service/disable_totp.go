package service

import (
	"github.com/go-playground/validator"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/mongo"
)

func DisableTotp(req model.DisableTotpReq) (model.DisableTotpRes, model.ServiceResp) {
	res := model.DisableTotpRes{}

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

	if totp.Status != string(config.Created) {
		return res, model.ServiceError.StatusConflictError("totp not already created")
	}

	err = mongo.GetInstance().DeleteTotp(req.AccountId, req.ServiceName)
	if err != nil {
		return res, model.ServiceError.InternalServiceError(err.Error())
	}

	return res, model.ServiceError.OK
}
