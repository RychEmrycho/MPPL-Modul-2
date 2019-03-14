package manage_product

import "github.com/jinzhu/gorm"

type Size struct {
	gorm.Model

	Name string

	// Size belongs to Category
	// db.Model(&user).Related(&profile)
	// SELECT * FROM profiles WHERE id = 111;
	// 111 is user's foreign key ProfileID
	Category   Category
	CategoryID uint
}
