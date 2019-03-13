package repository

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/jinzhu/gorm"
)

type brandRepository struct {
	Conn *gorm.DB
}

func NewBrandRepository(Conn *gorm.DB) manage_product.RepositoryBrand {
	return &brandRepository{Conn}
}

func (pr *brandRepository) Fetch() (res []*Brand, err error) {
	var brands []*Brand
	err = pr.Conn.Find(&brands).Error

	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (pr *brandRepository) GetById(id uint) (*Brand, error) {
	var brand_ Brand
	err := pr.Conn.Find(&brand_, id).Error

	if err != nil {
		return nil, err
	}

	return &brand_, nil
}

func (pr *brandRepository) Update(p *Brand) error {
	var brand_ Brand
	pr.Conn.Find(brand_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *brandRepository) Store(p *Brand) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *brandRepository) Delete(id uint) error {
	var brand_ Brand
	pr.Conn.Find(&brand_)
	err := pr.Conn.Delete(&brand_).Error

	if err != nil {
		return err
	}

	return nil
}