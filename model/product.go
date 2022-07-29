package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductID       int64  `gorm:"unique;not null"` // 产品id
	ProductName     string `gorm:"not null"`        // 产品名称
	ProductNum      uint64 `gorm:"not null"`        // 产品数量
	ProductImageUrl string `gorm:"not null"`        // 产品图片url
	ProductUrl      string `gorm:"not null"`        // 产品url
}
