package repository

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/jinzhu/gorm"
)

type categoryRepository struct {
	Conn *gorm.DB
}

func NewCategoryRepository(Conn *gorm.DB) manage_product.RepositoryCategory {
	return &categoryRepository{Conn}
}

func (pr *categoryRepository) Fetch() (res []*Category, err error) {
	var categorys []*Category
	err = pr.Conn.Find(&categorys).Error

	if err != nil {
		return nil, err
	}

	return categorys, nil
}

func (pr *categoryRepository) GetById(id uint) (*Category, error) {
	var category_ Category
	err := pr.Conn.Find(&category_, id).Error

	if err != nil {
		return nil, err
	}

	return &category_, nil
}

func (pr *categoryRepository) Update(p *Category) error {
	var category_ Category
	pr.Conn.Find(category_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *categoryRepository) Store(p *Category) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *categoryRepository) Delete(id uint) error {
	var category_ Category
	pr.Conn.Find(&category_)
	err := pr.Conn.Delete(&category_).Error

	if err != nil {
		return err
	}

	return nil
}