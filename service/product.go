package service

import (
	"iris-seckill/db/mysql"
	"iris-seckill/model"
)

type IProductService interface {
	GetProductByID(int64) (*model.Product, error)
	GetAllProduct() ([]*model.Product, error)
	DeleteProductByID(int64) error
	InsertProduct(product *model.Product) (int64, error)
	UpdateProduct(product *model.Product) error
}

type ProductService struct {
}

// NewProductService 初始化函数
func NewProductService() IProductService {
	return &ProductService{}
}

// GetProductByID 根据产品ID查询产品
func (p *ProductService) GetProductByID(productID int64) (*model.Product, error) {
	var product model.Product
	err := mysql.MysqlDB.Where("product_id = ?", productID).First(&product).Error()
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// GetAllProduct 查询所有产品
func (p *ProductService) GetAllProduct() (res []*model.Product, err error) {
	err = mysql.MysqlDB.Find(&res).Error()
	return
}

// DeleteProductByID 根据产品ID删除产品
func (p *ProductService) DeleteProductByID(productID int64) error {
	err := mysql.MysqlDB.Where("product_id = ?", productID).Delete(&model.Product{}).Error()
	return err
}

// InsertProduct 插入产品
func (p *ProductService) InsertProduct(product *model.Product) (int64, error) {
	err := mysql.MysqlDB.Create(product).Error()
	if err != nil {
		return 0, err
	}
	return product.ProductID, nil
}

// UpdateProduct 更新产品
func (p *ProductService) UpdateProduct(product *model.Product) error {
	err := mysql.MysqlDB.Save(product).Error()
	return err
}
