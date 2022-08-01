package model

type DisableTotpReq struct {
	AccountId   string `json:"accountId" validate:"required"`
	ServiceName string `json:"serviceName" validate:"required"`
}

type DisableTotpRes struct {
}
