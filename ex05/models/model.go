package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Product_Name string
	Price        int
}
type Cart struct {
	gorm.Model
	Products []*Product  `gorm:"many2many:cart_items"`
	CartItem []*CartItem // Many-to-many relationship with Product through CartItem
}
type CartItem struct {
	gorm.Model
	CartID    string `gorm:"primaryKey"`
	ProductID string `gorm:"primaryKey"`
	Quantity  int
}

type User struct {
	gorm.Model
	Email string `gorm:"unique"`
	Password string
}
