package manage_product

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model

	ImageURL  string
	RatePoint int
	Comment   string

	// Product has many reviews, ProductID is the foreign key
	ProductID uint
}
