package tests

import (
	"context"
	"github.com/aleynaguzell/challange-api/pkg/config"
	"github.com/aleynaguzell/challange-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

const testMongoUrl = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getir-case-study?retryWrites=true"

func init() {
	logger.Init()
	config.Setup("../")
}

func GetMongoClient(t *testing.T) (*mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(testMongoUrl))
	if err != nil {
		t.Fail()
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		t.Fail()
	}

	return client, nil
}
