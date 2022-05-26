package storage

import (
	"github.com/aleynaguzell/challange-api/storage/memory"
	"github.com/aleynaguzell/challange-api/storage/record"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	mongo  *MongoStorage
	memory *memory.Store
}

type MongoStorage struct {
	mongoClient *mongo.Client
	records     RecordRepository
}

func NewMongo(client *mongo.Client) *MongoStorage {
	return &MongoStorage{
		mongoClient: client,
		records:     record.NewRepository(client),
	}
}

func (d *Database) GetMemoryDb() *memory.Store {
	return d.memory
}

func (d *Database) GetMongoStorage() *MongoStorage {
	return d.mongo
}

func New(client *mongo.Client, memory *memory.Store) *Database {
	return &Database{
		mongo:  NewMongo(client),
		memory: memory,
	}
}

func (db *MongoStorage) GetRecordRepository() RecordRepository {
	return db.records
}
