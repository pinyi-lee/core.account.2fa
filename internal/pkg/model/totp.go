package model

type Totp struct {
	AccountId   string `bson:"accountId" json:"accountId"`
	ServiceName string `bson:"serviceName" json:"serviceName"`
	Secret      string `bson:"secret" json:"secret"`
	Status      string `bson:"status" json:"status"`
	CreatedAt   int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt   int64  `bson:"updatedAt" json:"updatedAt"`
}
