package model

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Records interface{} `json:"records"`
}

type Record struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int64     `json:"totalCount"`
}

func RespondWithJSON(w http.ResponseWriter, code int, payload Response) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
