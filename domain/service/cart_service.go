package service

import (
	"cart/domain/model"
	"cart/domain/repository"
)

type ICartService interface {
	InitTable() error
	AddCart(cart *model.Cart) (int64, error)
	DelCart(cartId int64) error
	UpdateCart(cart *model.Cart) error
	FindCartById(cartId int64) (*model.Cart, error)
	FindAllCart(userId int64) ([]*model.Cart, error)
	CleanCart(userId int64) error
	DecrNum(cartId int64, num int64) error
	IncrNum(cartId int64, num int64) error
}

func NewCartService(cartRepository repository.ICartRepository) ICartService {
	return &CartService{
		cartRepository,
	}
}

type CartService struct {
	repository repository.ICartRepository
}

func (c *CartService) InitTable() error {
	return c.repository.InitTable()
}

func (c *CartService) AddCart(cart *model.Cart) (int64, error) {
	return c.repository.AddCart(cart)
}

func (c *CartService) DelCart(cartId int64) error {
	return c.repository.DelCartById(cartId)
}

func (c *CartService) UpdateCart(cart *model.Cart) error {
	return c.repository.UpdateCart(cart)
}

func (c *CartService) FindCartById(cartId int64) (*model.Cart, error) {
	return c.repository.FindCartById(cartId)
}
func (c *CartService) FindAllCart(userId int64) ([]*model.Cart, error) {
	return c.repository.FindAll(userId)
}

func (c *CartService) CleanCart(userId int64) error {
	return c.repository.CleanCart(userId)
}

func (c *CartService) DecrNum(cartId int64, num int64) error {
	return c.repository.DecrNum(cartId, num)
}

func (c *CartService) IncrNum(cartId int64, num int64) error {
	return c.repository.IncrNum(cartId, num)
}
