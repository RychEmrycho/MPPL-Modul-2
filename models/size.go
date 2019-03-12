package models

import "github.com/jinzhu/gorm"

type Size struct {
	gorm.Model

	Name string
	Category Category
}