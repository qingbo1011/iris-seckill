package service

import (
	"iris-seckill/db/mysql"
	"iris-seckill/model"
)

type IOrderService interface {
	GetOrderByID(uint) (*model.Order, error)
	DeleteOrderByID(uint) error
	UpdateOrder(*model.Order) error
	InsertOrder(*model.Order) (uint, error)
	GetAllOrder() ([]*model.Order, error)
	GetAllOrderInfo() (map[int]map[string]string, error)
}

type OrderService struct {
}

func NewOrderService() IOrderService {
	return &OrderService{}
}

// GetOrderByID 根据ID查找订单
func (o *OrderService) GetOrderByID(orderID uint) (order *model.Order, err error) {
	err = mysql.MysqlDB.Where("id = ?", orderID).First(&order).Error
	return
}

// DeleteOrderByID 根据ID删除订单
func (o *OrderService) DeleteOrderByID(orderID uint) error {
	err := mysql.MysqlDB.Where("id = ?", orderID).Delete(&model.Order{}).Error
	return err
}

// UpdateOrder 更新订单
func (o *OrderService) UpdateOrder(order *model.Order) error {
	err := mysql.MysqlDB.Model(&model.Product{}).Where("id = ?", order.ID).Updates(order).Error
	return err
}

// InsertOrder 新增订单
func (o *OrderService) InsertOrder(order *model.Order) (uint, error) {
	err := mysql.MysqlDB.Create(order).Error
	return order.ID, err
}

// GetAllOrder 查询全部订单
func (o *OrderService) GetAllOrder() (res []*model.Order, err error) {
	err = mysql.MysqlDB.Find(&res).Error
	return
}

func (o *OrderService) GetAllOrderInfo() (map[int]map[string]string, error) {
	orderMap := make(map[int]map[string]string, 0)
	orders, err := o.GetAllOrder()
	if err != nil {
		return nil, err
	}
	for i, _ := range orders {
		orderMap[i] = map[string]string{"用户名order.UserID": "产品order.ProductID"}
	}
	return orderMap, nil
}
