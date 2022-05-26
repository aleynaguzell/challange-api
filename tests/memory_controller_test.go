package tests

import (
	"bytes"
	"context"
	"github.com/aleynaguzell/challange-api/controllers"
	"github.com/aleynaguzell/challange-api/storage"
	"github.com/aleynaguzell/challange-api/storage/memory"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInMemorySetController(t *testing.T) {

	expected := []byte(`{"key":"active-tabs","value":"getir"}`)

	postBody := []byte(`{"key":  "active-tabs","value": "getir"}`)
	req := httptest.NewRequest(http.MethodPost, "/in-memory/", bytes.NewBuffer(postBody))

	mStorage := memory.New()
	mClient, err := GetMongoClient(t)
	defer mClient.Disconnect(context.TODO())
	cf := controllers.NewControllerFactory(storage.New(mClient, mStorage))

	w := httptest.NewRecorder()
	cf.GetMemoryController().Set(w, req)

	res := w.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if http.StatusCreated != res.StatusCode {
		t.Error("expected ", http.StatusCreated, "got status", res.StatusCode)
	}
	if !strings.Contains(string(body), string(expected)) {
		t.Error("expected "+string(expected)+" got", string(body))
	}
}

func TestInMemoryGetController(t *testing.T) {

	postBody := []byte(`{"key":  "active-tabs","value": "getir"}`)
	reqPost := httptest.NewRequest(http.MethodPost, "/in-memory/", bytes.NewBuffer(postBody))

	mStorage := memory.New()
	mClient, err := GetMongoClient(t)
	defer mClient.Disconnect(context.TODO())
	cf := controllers.NewControllerFactory(storage.New(mClient, mStorage))

	wp := httptest.NewRecorder()
	cf.GetMemoryController().Set(wp, reqPost)

	expected := []byte(`{"key":"active-tabs","value":"getir"}`)
	req := httptest.NewRequest(http.MethodGet, "/in-memory?key=active-tabs", nil)

	w := httptest.NewRecorder()
	cf.GetMemoryController().Get(w, req)

	res := w.Result()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if http.StatusOK != res.StatusCode {
		t.Error("expected", http.StatusOK, "got status ", res.StatusCode)
	}
	if !strings.Contains(string(body), string(expected)) {
		t.Error("expected"+string(expected)+"got", string(body))
	}
}
