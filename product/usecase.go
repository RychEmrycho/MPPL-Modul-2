package product

import (
	"MPPL-Modul-2/models"
)

type UseCase interface {
	Fetch() (res []*models.Product, err error)
	GetById(id uint) (*models.Product, error)
	Update(p *models.Product) error
	Store(p *models.Product) error
	Delete(id uint) error
}