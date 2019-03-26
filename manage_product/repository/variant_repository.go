package repository

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/jinzhu/gorm"
)

type variantRepository struct {
	Conn *gorm.DB
}

func NewVariantRepository(Conn *gorm.DB) manage_product.RepositoryVariant {
	return &variantRepository{Conn}
}

func (pr *variantRepository) Fetch() (res []*Variant, err error) {
	var variants []*Variant
	err = pr.Conn.Find(&variants).Error

	if err != nil {
		return nil, err
	}

	return variants, nil
}

func (pr *variantRepository) GetById(id uint) (*Variant, error) {
	var variant_ Variant
	err := pr.Conn.Find(&variant_, id).Error

	if err != nil {
		return nil, err
	}

	return &variant_, nil
}

func (pr *variantRepository) Update(p *Variant, id uint) error {
	var variant_ Variant
	pr.Conn.Find(variant_, id)

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *variantRepository) Store(p *Variant) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *variantRepository) Delete(id uint) error {
	var variant_ Variant
	pr.Conn.Find(&variant_)
	err := pr.Conn.Delete(&variant_).Error

	if err != nil {
		return err
	}

	return nil
}