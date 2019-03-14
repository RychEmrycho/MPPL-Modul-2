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

	// Brand has many products, BrandID is the foreign key
	BrandID    uint
	// Category has many products, CategoryID is the foreign key
	CategoryID uint

	// Product has many reviews, images and variant
	// db.Model(&user).Related(&emails)
	// SELECT * FROM emails WHERE user_id = 111;
	// 111 is user's primary key
	Review  []Review
	Image   []Image
	Variant []Variant
}
