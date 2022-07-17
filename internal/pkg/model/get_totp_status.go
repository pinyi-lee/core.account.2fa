package model

type GetTotpStatusReq struct {
	AccountId   string `json:"accountId" validate:"required"`
	ServiceName string `json:"serviceName" validate:"required"`
}

type GetTotpStatusRes struct {
	Status string `json:"status"`
}
