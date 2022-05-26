package storage

import "github.com/aleynaguzell/getir-challange-api/model"

type RecordRepository interface {
	Get(req model.Request) ([]model.Record,error)
}