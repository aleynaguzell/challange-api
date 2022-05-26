package controllers

import (
	"encoding/json"
	"github.com/aleynaguzell/challange-api/model"
	"github.com/aleynaguzell/challange-api/storage"
	"net/http"
)

type RecordController struct {
	db *storage.MongoStorage
}

func NewRecordController(db *storage.MongoStorage) *RecordController {
	return &RecordController{
		db: db,
	}
}

//Fetch data from records collection
//HTTP.POST
func (m *RecordController) GetRecords(rw http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodPost {
		var result model.Response
		var recordQuery model.Request

		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&recordQuery)
		records, err := m.db.GetRecordRepository().Get(recordQuery)
		if err != nil {
			result.Code = http.StatusBadRequest
			result.Msg = err.Error()
			jData, _ := json.Marshal(result)

			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(jData)
			return
		}
		if len(records) == 0 {
			result.Msg = "Not Found"
		} else {
			result.Msg = "Success"
		}

		result.Code = 0
		result.Records = records
		jData, _ := json.Marshal(result)

		rw.WriteHeader(http.StatusOK)
		rw.Write(jData)
		return
	} else {
		http.Error(rw, "Method not found", http.StatusNotFound)
	}
}
