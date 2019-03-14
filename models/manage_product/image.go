package manage_product

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model

	ImageURL  string

	// Product has many images, ProductID is the foreign key
	ProductID uint
}
