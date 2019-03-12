package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model

	Name      string
	Price     int
	Weight    int
	Stock     int
	TotalView int
	Brand     Brand
	Category  Category
	Review    []Review
	Image     []Image
	Variant   []Variant
}
