package model

type Totp struct {
	AccountId   string `bson:"accountId" json:"accountId"`
	ServiceName string `bson:"serviceName" json:"serviceName"`
	Key         string `bson:"key" json:"key"`
	Status      string `bson:"status" json:"status"`
	CreatedAt   int64  `bson:"createdAt" json:"createdAt"`
	UpdatedAt   int64  `bson:"updatedAt" json:"updatedAt"`
}
