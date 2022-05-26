package mongo

import (
	"context"
	"fmt"
	"github.com/aleynaguzell/challange-api/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection = "records"
)

func Init() (*mongo.Client, error) {
	ctx := context.Background()
	fmt.Println("mongo conn reading",config.GetConfig().Mongo.Url )
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().Mongo.Url))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Mongo Connection Successful!")

	return client, nil
}
