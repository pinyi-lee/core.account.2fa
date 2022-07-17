package service

import (
	"github.com/go-playground/validator"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
)

func VerifyTotp(req model.VerifyTotpReq) (model.VerifyTotpRes, model.ServiceResp) {
	res := model.VerifyTotpRes{}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return res, model.ServiceError.BadRequestError(err.Error())
	}

	return res, model.ServiceError.OK
}
