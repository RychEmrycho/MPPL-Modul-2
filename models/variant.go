package models

import "github.com/jinzhu/gorm"

type Variant struct {
	gorm.Model

	Product Product
	Color Color
	Size Size
}