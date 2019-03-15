package manage_product

import "github.com/jinzhu/gorm"

type Variant struct {
	gorm.Model

	// Variant has one color and one size
	// var card CreditCard
	// db.Model(&user).Related(&card, "CreditCard")
	// SELECT * FROM credit_cards WHERE user_id = 123; // 123 is user's primary key
	// CreditCard is user's field name, it means get user's CreditCard relations and fill it into variable card
	// If the field name is same as the variable's type name, like above example, it could be omitted, like:
	// db.Model(&user).Related(&card)
	Color   Color
	ColorID uint `json:"-"`
	Size    Size
	SizeID  uint `json:"-"`

	// Product has many variants, ProductID is the foreign key
	ProductID uint `json:"-"`
}
