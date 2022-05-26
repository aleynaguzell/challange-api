package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/aleynaguzell/challange-api/model"
	"github.com/aleynaguzell/challange-api/pkg/logger"
	"github.com/aleynaguzell/challange-api/storage/memory"
	"net/http"
)

type MemoryController struct {
	db *memory.Store
}

func NewMemoryController(db *memory.Store) *MemoryController {
	return &MemoryController{
		db: db,
	}
}

//Set data to an in-memory database
//HTTP.POST
func (c *MemoryController) Set(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var request model.MemoryReq
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			_, err := fmt.Fprintf(w, "%+v", err.Error())
			if err != nil {
				logger.Logger.Error("Request can not decoded", err)
				return
			}
		}

		_ = c.db.Store(request.Key, request.Value)
		w.WriteHeader(http.StatusCreated)

		err = json.NewEncoder(w).Encode(&request)
		if err != nil {
			logger.Logger.Error("Request can not encoded", err)
			return
		}
	} else {
		http.Error(w, "Method not found", http.StatusNotFound)
	}
}

//Fetch data from an in-memory database
//HTTP.GET
func (c *MemoryController) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		keyQuery := r.URL.Query()
		key := keyQuery.Get("key")
		value, err := c.db.Get(key)
		if err != nil {
			_, err := fmt.Fprintf(w, "%+v", err.Error())
			if err != nil {
				logger.Logger.Error(err)
				return
			}
		} else {
			out := model.MemoryReq{Key: key, Value: value}
			err = json.NewEncoder(w).Encode(out)
			if err != nil {
				logger.Logger.Error(err)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
	} else {
		http.Error(w, "Method not found", http.StatusNotFound)
	}
}
