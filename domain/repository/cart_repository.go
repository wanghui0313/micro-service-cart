package repository

import (
	"cart/domain/model"
	"gorm.io/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartById(int64) (*model.Cart, error)
	AddCart(*model.Cart) (int64, error)
	DelCartById(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]*model.Cart, error)
	CleanCart(int642 int64) error
	IncrNum(int64, int64) error
	DecrNum(int64, int64) error
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{
		mysqlDb: db,
	}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}

func (c *CartRepository) InitTable() error {
	return c.mysqlDb.AutoMigrate(&model.Cart{})
}

func (c *CartRepository) FindCartById(id int64) (cart *model.Cart, err error) {
	return cart, c.mysqlDb.First(cart, id).Error
}

func (c *CartRepository) AddCart(cart *model.Cart) (id int64, err error) {
	var findCart *model.Cart
	c.mysqlDb.Where("productId=? and sizeId=? and userId=?", cart.ProductId, cart.SizeId, cart.UserId).First(findCart)
	if findCart.Id != 0 {
		return findCart.Id, c.mysqlDb.Update("num", gorm.Expr("num+1")).Error
	} else {
		return cart.Id, c.mysqlDb.Create(cart).Error
	}
}

func (c *CartRepository) DelCartById(id int64) error {
	return c.mysqlDb.Where("id=?", id).Delete(&model.Cart{}).Error
}

func (c *CartRepository) UpdateCart(cart *model.Cart) error {
	return c.mysqlDb.Save(cart).Error
}

func (c *CartRepository) FindAll(userId int64) (carts []*model.Cart, err error) {
	return carts, c.mysqlDb.Where("user_id=?", userId).Find(carts).Error
}

func (c *CartRepository) CleanCart(userId int64) error {
	return c.mysqlDb.Where("user_id=?", userId).Delete(&model.Cart{}).Error
}

func (c *CartRepository) IncrNum(id int64, num int64) error {
	return c.mysqlDb.Model(&model.Cart{Id: id}).Update("num", gorm.Expr("num+?", num)).Error
}

func (c *CartRepository) DecrNum(id int64, num int64) error {
	return c.mysqlDb.Model(&model.Cart{Id: id}).Update("num", gorm.Expr("num-?", num)).Error
}
