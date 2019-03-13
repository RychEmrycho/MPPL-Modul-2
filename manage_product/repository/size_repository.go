package repository

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/jinzhu/gorm"
)

type sizeRepository struct {
	Conn *gorm.DB
}

func NewSizeRepository(Conn *gorm.DB) manage_product.RepositorySize {
	return &sizeRepository{Conn}
}

func (pr *sizeRepository) Fetch() (res []*Size, err error) {
	var sizes []*Size
	err = pr.Conn.Find(&sizes).Error

	if err != nil {
		return nil, err
	}

	return sizes, nil
}

func (pr *sizeRepository) GetById(id uint) (*Size, error) {
	var size_ Size
	err := pr.Conn.Find(&size_, id).Error

	if err != nil {
		return nil, err
	}

	return &size_, nil
}

func (pr *sizeRepository) Update(p *Size) error {
	var size_ Size
	pr.Conn.Find(size_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *sizeRepository) Store(p *Size) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *sizeRepository) Delete(id uint) error {
	var size_ Size
	pr.Conn.Find(&size_)
	err := pr.Conn.Delete(&size_).Error

	if err != nil {
		return err
	}

	return nil
}