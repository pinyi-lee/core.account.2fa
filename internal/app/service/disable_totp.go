package service

import (
	"github.com/go-playground/validator"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
)

func DisableTotp(req model.DisableTotpReq) (model.DisableTotpRes, model.ServiceResp) {
	res := model.DisableTotpRes{}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return res, model.ServiceError.BadRequestError(err.Error())
	}

	return res, model.ServiceError.OK
}
