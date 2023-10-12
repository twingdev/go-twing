package core

import "github.com/syndtr/goleveldb/leveldb"

type IDatabase interface {
	DB() *leveldb.DB
}

type datastore struct{}

func (ds *datastore) DB() *leveldb.DB {
	db, err := leveldb.OpenFile("db/db", nil)
	if err != nil {
		panic(err)
	}
	return db
}

func (ds *datastore) Get() {
	defer ds.DB().Close()
}

func NewDatastore() IDatabase {
	return &datastore{}
}
