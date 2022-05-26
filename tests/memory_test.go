package tests

import (
	"github.com/aleynaguzell/challange-api/storage/memory"
	"testing"
)

var Store *memory.Store //declare

func init() {
	Store = memory.New() //assign
}

func TestStore(t *testing.T) {

	nilResult := Store.Store("active-tabs", "getir")
	if nilResult != nil {
		t.Fail()
	}
	if value, err := Store.Get("active-tabs"); err != nil || value != "getir" {
		t.Fail()
	}
}

func TestGet(t *testing.T) {

	nilResult := Store.Store("active-tabs", "getir")
	if nilResult != nil {
		t.Fail()
	}

	value, err := Store.Get("active-tabs")
	if err != nil || value != "getir" {
		t.Fail()
	}
}
