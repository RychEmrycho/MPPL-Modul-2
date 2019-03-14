package manage_product

import "github.com/jinzhu/gorm"

type Brand struct {
	gorm.Model

	Name string

	// Brand has many products
	//db.Model(&user).Related(&emails)
	// SELECT * FROM emails WHERE user_id = 111; // 111 is user's primary key
	Products []Product
}
