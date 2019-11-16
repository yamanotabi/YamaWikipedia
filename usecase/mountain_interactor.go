package usecase

import (
	"github.com/shiki-tak/YamaWikipedia/domain"
)

type MountainInteractor struct {
	MountainRepository MountainRepository
}

func (interactor *MountainInteractor) Add(m domain.Mountain) error {
	return interactor.MountainRepository.Store(m)
}

func (interactor *MountainInteractor) AllMountains() (domain.Mountains, error) {
	mountains, err := interactor.MountainRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return mountains, nil
}

func (interactor *MountainInteractor) MountainById(key string) (domain.Mountain, error) {
	mountain, err := interactor.MountainRepository.FindById(key)
	if err != nil {
		return domain.Mountain{}, err
	}
	return mountain, nil
}
