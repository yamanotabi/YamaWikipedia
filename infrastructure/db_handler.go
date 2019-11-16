package infrastructure

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

type LevelDBHandler struct {
	Conn *leveldb.DB
}

func NewLevelDBHandler() *LevelDBHandler {
	// open db
	conn, err := leveldb.OpenFile("./db", nil)
	if err != nil {
		panic(err.Error)
	}
	levelDBHandler := new(LevelDBHandler)
	levelDBHandler.Conn = conn

	return levelDBHandler
}

func (handler *LevelDBHandler) Put(key []byte, value []byte, args ...interface{}) error {
	err := handler.Conn.Put(key, value, nil)
	if err != nil {
		return fmt.Errorf("Put Data error:%s", err)
	}
	return nil
}

func (handler *LevelDBHandler) Get(key string, args ...interface{}) ([]byte, error) {
	jsonBytes, err := handler.Conn.Get([]byte(key), nil)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func (handler *LevelDBHandler) Scan() iterator.Iterator {
	return handler.Conn.NewIterator(nil, nil)
}
