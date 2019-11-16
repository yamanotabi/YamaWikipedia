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

func (repo *MountainRepository) FindAll() (domain.Mountains, error) {
	ｍountains := []domain.Mountain{}
	iter := repo.Scan()
	for iter.Next() {
		m := new(domain.Mountain)
		value := iter.Value()
		err := json.Unmarshal(value, m)
		if err != nil {
			return nil, err
		}

		ｍountains = append(ｍountains, *m)
	}

	return ｍountains, nil
}
