package service

import (
	"github.com/go-playground/validator"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/config"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/mongo"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/util"
)

func InitTotp(req model.InitTotpReq) (model.InitTotpRes, model.ServiceResp) {
	res := model.InitTotpRes{}

	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return res, model.ServiceError.BadRequestError(err.Error())
	}

	totp, err := mongo.GetInstance().GetTotp(req.AccountId, req.ServiceName)
	if err != nil && err != mongo.ERROR_DATA_NOT_FOUND {
		return res, model.ServiceError.InternalServiceError(err.Error())
	}

	if totp.Status == string(config.Created) {
		return res, model.ServiceError.StatusConflictError("totp already created")
	}

	key, err := GenerateKey(req.AccountId, req.ServiceName)
	if err != nil {
		return res, model.ServiceError.InternalServiceError(err.Error())
	}

	qrCode, err := GetQrCode(key)
	if err != nil {
		return res, model.ServiceError.InternalServiceError(err.Error())
	}

	createReq := model.Totp{
		AccountId:   req.AccountId,
		ServiceName: req.ServiceName,
		Secret:      key.Secret(),
		Status:      string(config.Init),
		CreatedAt:   util.Timestamp(),
		UpdatedAt:   util.Timestamp(),
	}
	err = mongo.GetInstance().CreateTotp(createReq)
	if err != nil {
		return res, model.ServiceError.InternalServiceError(err.Error())
	}

	res.QRCode = qrCode
	res.Secret = key.Secret()
	return res, model.ServiceError.OK
}
