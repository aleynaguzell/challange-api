package controllers

import "github.com/aleynaguzell/getir-challange-api/storage"

type ControllerFactory struct {
	db *storage.Database
}

func NewControllerFactory(db *storage.Database) *ControllerFactory{
	return  &ControllerFactory{
		db: db,
	}
}

func (c *ControllerFactory) GetMemoryController() *MemoryController{
	return NewMemoryController(c.db.GetMemoryDb())
}

func (c *ControllerFactory) GetRecordController() *RecordController{
	return NewRecordController(c.db.GetMongoStorage())
}