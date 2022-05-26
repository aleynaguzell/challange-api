package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/aleynaguzell/getir-challange-api/model"
	"github.com/aleynaguzell/getir-challange-api/pkg/logger"
	"github.com/aleynaguzell/getir-challange-api/storage/memory"
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

func (c *MemoryController) Set(w http.ResponseWriter, r *http.Request) {
	var input model.MemoryReq

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		_, err := fmt.Fprintf(w, "%+v", err.Error())
		if err != nil {
			logger.Logger.Error("Request can not decoded", err)
			return
		}
	}

	_ = c.db.Store(input.Key, input.Value)
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(&input)
	if err != nil {
		logger.Logger.Error("Request can not encoded", err)
		return
	}
}

func (c *MemoryController) Get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query()
	keyQuery := key.Get("key")
	value, err := c.db.Get(keyQuery)
	if err != nil {
		_, err := fmt.Fprintf(w, "%+v", err.Error())
		if err != nil {
			logger.Logger.Error(err)
			return
		}
	} else {
		out := model.MemoryReq{Key: keyQuery, Value: value}
		err = json.NewEncoder(w).Encode(out)
		if err != nil {
			logger.Logger.Error(err)
			return
		}
	}
}
