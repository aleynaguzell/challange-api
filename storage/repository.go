package storage

import "github.com/aleynaguzell/challange-api/model"

type RecordRepository interface {
	Get(req model.Request) ([]model.Record, error)
}
