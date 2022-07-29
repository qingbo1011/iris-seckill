package request

type ProductReq struct {
	ProductID       int64  `json:"productID" form:"productID"`             // 产品id
	ProductName     string `json:"productName" form:"productName"`         // 产品名称
	ProductNum      uint64 `json:"productNum" form:"productNum"`           // 产品数量
	ProductImageUrl string `json:"productImageUrl" form:"productImageUrl"` // 产品图片url
	ProductUrl      string `json:"productUrl" form:"productUrl"`           // 产品url
}
