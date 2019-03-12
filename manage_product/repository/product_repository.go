package repository

import (
	. "MPPL-Modul-2/models/manage_product"
	"MPPL-Modul-2/manage_product"
	"github.com/jinzhu/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) manage_product.Repository {
	return &productRepository{Conn}
}

func (pr *productRepository) Fetch() (res []*Product, err error) {
	var products []*Product
	err = pr.Conn.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *productRepository) GetById(id uint) (*Product, error) {
	var product_ Product
	err := pr.Conn.Find(&product_, id).Error

	if err != nil {
		return nil, err
	}

	return &product_, nil
}

func (pr *productRepository) Update(p *Product) error {
	var product_ Product
	pr.Conn.Find(product_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *productRepository) Store(p *Product) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *productRepository) Delete(id uint) error {
	var product_ Product
	pr.Conn.Find(&product_)
	err := pr.Conn.Delete(&product_).Error

	if err != nil {
		return err
	}

	return nil
}