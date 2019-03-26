package repository

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/jinzhu/gorm"
)

type colorRepository struct {
	Conn *gorm.DB
}

func NewColorRepository(Conn *gorm.DB) manage_product.RepositoryColor {
	return &colorRepository{Conn}
}

func (pr *colorRepository) Fetch() (res []*Color, err error) {
	var colors []*Color
	err = pr.Conn.Find(&colors).Error

	if err != nil {
		return nil, err
	}

	return colors, nil
}

func (pr *colorRepository) GetById(id uint) (*Color, error) {
	var color_ Color
	err := pr.Conn.Find(&color_, id).Error

	if err != nil {
		return nil, err
	}

	return &color_, nil
}

func (pr *colorRepository) Update(p *Color, id uint) error {
	var color_ Color
	pr.Conn.Find(color_, id)

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *colorRepository) Store(p *Color) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *colorRepository) Delete(id uint) error {
	var color_ Color
	pr.Conn.Find(&color_)
	err := pr.Conn.Delete(&color_).Error

	if err != nil {
		return err
	}

	return nil
}