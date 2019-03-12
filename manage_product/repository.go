package manage_product

import (
	. "MPPL-Modul-2/models/manage_product"
)

type Repository interface {
	Fetch() (res []*Product, err error)
	GetById(id uint) (*Product, error)
	Update(p *Product) error
	Store(p *Product) error
	Delete(id uint) error
}
