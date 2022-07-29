package service

import (
	"iris-seckill/db/mysql"
	"iris-seckill/model"
)

type IProductService interface {
	GetProductByID(int64) (*model.Product, error)
	GetAllProduct() ([]*model.Product, error)
	DeleteProductByID(int64) error
	InsertProduct(product *model.Product) (uint, error)
	UpdateProduct(product *model.Product) error
}

type ProductService struct {
}

// NewProductService 初始化函数
func NewProductService() IProductService {
	return &ProductService{}
}

// GetProductByID 根据产品ID查询产品
func (p *ProductService) GetProductByID(id int64) (*model.Product, error) {
	var product model.Product
	err := mysql.MysqlDB.Where("id = ?", id).First(&product).Error
	return &product, err
}

// GetAllProduct 查询所有产品
func (p *ProductService) GetAllProduct() (res []*model.Product, err error) {
	err = mysql.MysqlDB.Find(&res).Error
	return
}

// DeleteProductByID 根据产品ID删除产品
func (p *ProductService) DeleteProductByID(id int64) error {
	err := mysql.MysqlDB.Where("id = ?", id).Delete(&model.Product{}).Error
	return err
}

// InsertProduct 插入产品
func (p *ProductService) InsertProduct(product *model.Product) (uint, error) {
	err := mysql.MysqlDB.Create(product).Error
	return product.ID, err
}

// UpdateProduct 更新产品（注意下面的id，到时候写的时候要修改一下）
func (p *ProductService) UpdateProduct(product *model.Product) error {
	err := mysql.MysqlDB.Model(&model.Product{}).Where("id = ?", 2).Updates(product).Error
	return err
}
