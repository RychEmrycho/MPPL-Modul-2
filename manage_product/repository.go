package manage_product

import (
	. "MPPL-Modul-2/models/manage_product"
)

type RepositoryProduct interface {
	Fetch() (res []*Product, err error)
	GetById(id uint) (*Product, error)
	Update(p *Product) error
	Store(p *Product) error
	Delete(id uint) error
}

type RepositoryBrand interface {
	Fetch() (res []*Brand, err error)
	GetById(id uint) (*Brand, error)
	Update(p *Brand) error
	Store(p *Brand) error
	Delete(id uint) error
}

type RepositoryCategory interface {
	Fetch() (res []*Category, err error)
	GetById(id uint) (*Category, error)
	Update(p *Category) error
	Store(p *Category) error
	Delete(id uint) error
}

type RepositoryColor interface {
	Fetch() (res []*Color, err error)
	GetById(id uint) (*Color, error)
	Update(p *Color) error
	Store(p *Color) error
	Delete(id uint) error
}

type RepositoryImage interface {
	Fetch() (res []*Image, err error)
	GetById(id uint) (*Image, error)
	Update(p *Image) error
	Store(p *Image) error
	Delete(id uint) error
}

type RepositoryReview interface {
	Fetch() (res []*Review, err error)
	GetById(id uint) (*Review, error)
	Update(p *Review) error
	Store(p *Review) error
	Delete(id uint) error
}

type RepositorySize interface {
	Fetch() (res []*Size, err error)
	GetById(id uint) (*Size, error)
	Update(p *Size) error
	Store(p *Size) error
	Delete(id uint) error
}

type RepositoryVariant interface {
	Fetch() (res []*Variant, err error)
	GetById(id uint) (*Variant, error)
	Update(p *Variant) error
	Store(p *Variant) error
	Delete(id uint) error
}