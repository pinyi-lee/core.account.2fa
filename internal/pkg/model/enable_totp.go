package model

type EnableTotpReq struct {
	AccountId   string `json:"accountId" validate:"required"`
	ServiceName string `json:"serviceName" validate:"required"`
	Passcode    string `json:"passcode" validate:"required"`
}

type EnableTotpRes struct {
}
