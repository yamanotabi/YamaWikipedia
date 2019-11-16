package database

type LevelDBHandler interface {
	Put([]byte, []byte, ...interface{}) error
	Get(string, ...interface{}) ([]byte, error)
	Scan() [][]byte
}
