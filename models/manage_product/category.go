package manage_product

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model

	Name string
}