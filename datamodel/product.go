package datamodel

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductID       int64  `json:"productID"`       // 产品id
	ProductName     string `json:"productName"`     // 产品名称
	ProductNum      uint64 `json:"productNum"`      // 产品数量
	ProductImageUrl string `json:"productImageUrl"` // 产品图片url
	ProductUrl      string `json:"productUrl"`      // 产品url
}
