package manage_product

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model

	ImageURL  string
	ProductID uint
}
