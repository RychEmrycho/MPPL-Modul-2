package repository

import (
	"MPPL-Modul-2/models"
	"MPPL-Modul-2/product"
	"github.com/jinzhu/gorm"
)

type productRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(Conn *gorm.DB) product.Repository {
	return &productRepository{Conn}
}

func (pr *productRepository) Fetch() (res []*models.Product, err error) {
	var products []*models.Product
	err = pr.Conn.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (pr *productRepository) GetById(id uint) (*models.Product, error) {
	var product_ models.Product
	err := pr.Conn.Find(&product_, id).Error

	if err != nil {
		return nil, err
	}

	return &product_, nil
}

func (pr *productRepository) Update(p *models.Product) error {
	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *productRepository) Store(p *models.Product) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *productRepository) Delete(id uint) error {
	var product_ models.Product
	pr.Conn.Find(product_)
	err := pr.Conn.Delete(&product_).Error

	if err != nil {
		return err
	}

	return nil
}
