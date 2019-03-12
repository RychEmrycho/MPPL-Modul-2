package manage_product

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

	BrandID    uint
	CategoryID uint

	Review     []Review
	Image      []Image
	Variant    []Variant
}
