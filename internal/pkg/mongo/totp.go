package mongo

import (
	"context"

	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/model"
	"github.com/pinyi-lee/core.account.2fa.git/internal/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collectionTotp = "totp"
)

func (manager *Manager) CreateTotp(totp model.Totp) (err error) {
	collection := manager.client.Database(manager.databaseName).Collection(collectionTotp)

	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"accountId":   totp.AccountId,
		"serviceName": totp.ServiceName,
	}
	update := bson.D{{"$set", totp}}

	_, err = collection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func (manager *Manager) GetTotp(accountId string, serviceName string) (totp model.Totp, err error) {
	collection := manager.client.Database(manager.databaseName).Collection(collectionTotp)

	filter := bson.M{
		"accountId":   accountId,
		"serviceName": serviceName,
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&totp)
	if err == mongo.ErrNoDocuments {
		return totp, ERROR_DATA_NOT_FOUND
	}

	return totp, err
}

func (manager *Manager) UpdateTotp(accountId string, serviceName string, status string) error {
	collection := manager.client.Database(manager.databaseName).Collection(collectionTotp)

	filter := bson.M{
		"accountId":   accountId,
		"serviceName": serviceName,
	}

	update := bson.M{"$set": bson.M{"status": status, "createdAt": util.Timestamp()}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (manager *Manager) DeleteTotp(accountId string, serviceName string) error {
	collection := manager.client.Database(manager.databaseName).Collection(collectionTotp)

	filter := bson.M{
		"accountId":   accountId,
		"serviceName": serviceName,
	}

	_, err := collection.DeleteOne(context.TODO(), filter)
	return err
}
