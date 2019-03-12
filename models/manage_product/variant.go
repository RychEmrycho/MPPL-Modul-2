package manage_product

import "github.com/jinzhu/gorm"

type Variant struct {
	gorm.Model

	ColorID   uint
	SizeID    uint
	ProductID uint
}
