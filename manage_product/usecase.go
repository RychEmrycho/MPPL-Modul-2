package manage_product

import (
	. "MPPL-Modul-2/models/manage_product"
)

type UseCase interface {
	FetchProduct() (res []*Product, err error)
	GetByIdProduct(id uint) (*Product, error)
	UpdateProduct(p *Product) error
	StoreProduct(p *Product) error
	DeleteProduct(id uint) error

	FetchBrand() (res []*Brand, err error)
	GetByIdBrand(id uint) (*Brand, error)
	UpdateBrand(p *Brand) error
	StoreBrand(p *Brand) error
	DeleteBrand(id uint) error

	FetchCategory() (res []*Category, err error)
	GetByIdCategory(id uint) (*Category, error)
	UpdateCategory(p *Category) error
	StoreCategory(p *Category) error
	DeleteCategory(id uint) error

	FetchColor() (res []*Color, err error)
	GetByIdColor(id uint) (*Color, error)
	UpdateColor(p *Color) error
	StoreColor(p *Color) error
	DeleteColor(id uint) error

	FetchImage() (res []*Image, err error)
	GetByIdImage(id uint) (*Image, error)
	UpdateImage(p *Image) error
	StoreImage(p *Image) error
	DeleteImage(id uint) error

	FetchReview() (res []*Review, err error)
	GetByIdReview(id uint) (*Review, error)
	UpdateReview(p *Review) error
	StoreReview(p *Review) error
	DeleteReview(id uint) error

	FetchSize() (res []*Size, err error)
	GetByIdSize(id uint) (*Size, error)
	UpdateSize(p *Size) error
	StoreSize(p *Size) error
	DeleteSize(id uint) error

	FetchVariant() (res []*Variant, err error)
	GetByIdVariant(id uint) (*Variant, error)
	UpdateVariant(p *Variant) error
	StoreVariant(p *Variant) error
	DeleteVariant(id uint) error
}