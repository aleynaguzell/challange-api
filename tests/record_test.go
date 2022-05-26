package tests

import (
	"context"
	"github.com/aleynaguzell/challange-api/model"
	"github.com/aleynaguzell/challange-api/storage/record"
	"testing"
)

func TestGetRecord(t *testing.T) {
	mClient, err := GetMongoClient(t)
	defer mClient.Disconnect(context.TODO())
	if err != nil {
		t.Fail()
	}
	repo := record.NewRepository(mClient) //assign

	recordQuery := model.Request{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  1900,
		MaxCount:  2700,
	}

	records, err := repo.Get(recordQuery)
	if err != nil {
		t.Fail()
	} else if len(records) == 0 {
		t.Fail()
	}
}
