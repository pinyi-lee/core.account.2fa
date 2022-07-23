package model

type InitTotpReq struct {
	AccountId   string `json:"accountId" validate:"required"`
	ServiceName string `json:"serviceName" validate:"required"`
}

type InitTotpRes struct {
	QRCode string `json:"qrCode"`
	Secret string `json:"secret"`
}
