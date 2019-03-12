package models

import "github.com/jinzhu/gorm"

type Review struct {
	gorm.Model

	ImageURL string
	RatePoint int
	Comment string

	Product Product
}