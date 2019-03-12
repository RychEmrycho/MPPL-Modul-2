package manage_product

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model

	ImageURL  string
	RatePoint int
	Comment   string

	ProductID uint
}
