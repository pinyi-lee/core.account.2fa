package service

import (
	"github.com/go-playground/validator"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
)

func InitTotp(req model.InitTotpReq) (model.InitTotpRes, model.ServiceResp) {
	res := model.InitTotpRes{}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return res, model.ServiceError.BadRequestError(err.Error())
	}

	return res, model.ServiceError.OK
}
