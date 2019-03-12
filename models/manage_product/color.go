package manage_product

import "github.com/jinzhu/gorm"

type Color struct {
	gorm.Model

	Name string
}