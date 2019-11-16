package usecase

import "github.com/shiki-tak/YamaWikipedia/domain"

type MountainRepository interface {
	Store(domain.Mountain) error
	FindById(string) (domain.Mountain, error)
	FindAll() (domain.Mountains, error)
}
