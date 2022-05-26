package record

import (
	"context"
	"github.com/aleynaguzell/challange-api/model"
	"github.com/aleynaguzell/challange-api/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository struct {
	client *mongo.Client
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		client: client,
	}
}

const collection = "records"

func (r *Repository) Get(req model.Request) ([]model.Record, error) {

	var results []model.Record

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {

		return nil, err
	}

	filter := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": startDate,
					"$lt": endDate,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": req.MinCount,
					"$lt": req.MaxCount,
				},
			},
		},
	}

	cursor, err := r.client.Database(config.GetConfig().Mongo.Database).Collection(collection).Aggregate(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var record model.Record
		err := cursor.Decode(&record)
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}

	return results, nil
}
