package repository

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/jinzhu/gorm"
)

type imageRepository struct {
	Conn *gorm.DB
}

func NewImageRepository(Conn *gorm.DB) manage_product.RepositoryImage {
	return &imageRepository{Conn}
}

func (pr *imageRepository) Fetch() (res []*Image, err error) {
	var images []*Image
	err = pr.Conn.Find(&images).Error

	if err != nil {
		return nil, err
	}

	return images, nil
}

func (pr *imageRepository) GetById(id uint) (*Image, error) {
	var image_ Image
	err := pr.Conn.Find(&image_, id).Error

	if err != nil {
		return nil, err
	}

	return &image_, nil
}

func (pr *imageRepository) Update(p *Image) error {
	var image_ Image
	pr.Conn.Find(image_, )

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *imageRepository) Store(p *Image) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *imageRepository) Delete(id uint) error {
	var image_ Image
	pr.Conn.Find(&image_)
	err := pr.Conn.Delete(&image_).Error

	if err != nil {
		return err
	}

	return nil
}