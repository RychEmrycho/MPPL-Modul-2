package manage_product

import "github.com/jinzhu/gorm"

type Size struct {
	gorm.Model

	Name       string
	CategoryID uint
}
