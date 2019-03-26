package repository

import (
	"MPPL-Modul-2/manage_product"
	. "MPPL-Modul-2/models/manage_product"
	"github.com/jinzhu/gorm"
)

type reviewRepository struct {
	Conn *gorm.DB
}

func NewReviewRepository(Conn *gorm.DB) manage_product.RepositoryReview {
	return &reviewRepository{Conn}
}

func (pr *reviewRepository) Fetch() (res []*Review, err error) {
	var reviews []*Review
	err = pr.Conn.Find(&reviews).Error

	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (pr *reviewRepository) GetById(id uint) (*Review, error) {
	var review_ Review
	err := pr.Conn.Find(&review_, id).Error

	if err != nil {
		return nil, err
	}

	return &review_, nil
}

func (pr *reviewRepository) Update(p *Review,id uint) error {
	var review_ Review
	pr.Conn.Find(review_, id)

	err := pr.Conn.Save(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *reviewRepository) Store(p *Review) error {
	err := pr.Conn.Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func (pr *reviewRepository) Delete(id uint) error {
	var review_ Review
	pr.Conn.Find(&review_)
	err := pr.Conn.Delete(&review_).Error

	if err != nil {
		return err
	}

	return nil
}