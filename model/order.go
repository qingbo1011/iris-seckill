package model

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	//OrderID     uint
	UserID      uint
	ProductID   uint
	OrderStatus int8
}

const (
	OrderWait    = iota // 0
	OrderSuccess        // 1
	OrderFailed         // 2
)
