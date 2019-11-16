package database

import "github.com/syndtr/goleveldb/leveldb/iterator"

type LevelDBHandler interface {
	Put([]byte, []byte, ...interface{}) error
	Get(string, ...interface{}) ([]byte, error)
	Scan() iterator.Iterator
}
