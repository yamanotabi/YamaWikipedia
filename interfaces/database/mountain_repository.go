package database

import (
	"encoding/json"
	"strconv"

	"github.com/shiki-tak/YamaWikipedia/domain"
)

type MountainRepository struct {
	LevelDBHandler
}

var (
	seq = 0
)

func (repo *MountainRepository) Store(m domain.Mountain) error {
	m.ID = seq
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	seqStr := strconv.Itoa(seq)
	err = repo.Put([]byte(seqStr), jsonBytes)

	if err != nil {
		return err
	}

	seq++
	return nil
}

func (repo *MountainRepository) FindById(key string) (domain.Mountain, error) {
	resAsBytes, err := repo.Get(key)
	if err != nil {
		return domain.Mountain{}, err
	}
	mountain := new(domain.Mountain)
	err = json.Unmarshal(resAsBytes, mountain)
	if err != nil {
		return domain.Mountain{}, err
	}

	return *mountain, nil
}

// TODO: implement
func (repo *MountainRepository) FindAll() (domain.Mountains, error) {
	return domain.Mountains{}, nil
}
