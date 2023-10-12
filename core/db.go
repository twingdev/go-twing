package core

import "github.com/syndtr/goleveldb/leveldb"

type IDatabase interface {
	DB() *leveldb.DB
}
