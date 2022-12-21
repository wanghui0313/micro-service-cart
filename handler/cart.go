package handler

import (
	"cart/domain/model"
	"cart/domain/service"
	pb "cart/proto"
	"context"
	"github.com/wanghui0313/micro-service-common/common"
)

type Cart struct {
	dataService service.ICartService
}

func (c *Cart) Add(ctx context.Context, info *pb.CartInfo, add *pb.ResAdd) error {
	var cart *model.Cart
	err := common.SwapTo(info, cart)
	if err != nil {
		return err
	}
	addId, err := c.dataService.AddCart(cart)
	if err != nil {
		return err
	}
	add.CartId = addId
	add.Msg = "添加成功"
	return nil
}

func (c *Cart) Clear(ctx context.Context, clean *pb.Clean, res *pb.Res) error {
	if err := c.dataService.CleanCart(clean.UserId); err != nil {
		return err
	}
	res.Msg = "清空购物车成功"
	return nil
}

func (c *Cart) Incr(ctx context.Context, item *pb.Item, res *pb.Res) error {
	if err := c.dataService.IncrNum(item.Id, item.ChangeNum); err != nil {
		return err
	}
	res.Msg = "购物车添加成功"
	return nil
}

func (c *Cart) Decr(ctx context.Context, item *pb.Item, res *pb.Res) error {
	if err := c.dataService.DecrNum(item.Id, item.ChangeNum); err != nil {
		return err
	}
	res.Msg = "购物车减少成功"
	return nil
}

func (c *Cart) DelItemById(ctx context.Context, id *pb.CartId, res *pb.Res) error {
	if err := c.dataService.DelCart(id.Id); err != nil {
		return err
	}
	res.Msg = "购物车删除成功"
	return nil
}

func (c *Cart) GetAll(ctx context.Context, all *pb.CartFindAll, all2 *pb.CartAll) error {
	carts, err := c.dataService.FindAllCart(all.UserId)
	if err != nil {
		return err
	}
	for _, cart := range carts {
		cartInfo := &pb.CartInfo{}
		err = common.SwapTo(cart, cartInfo)
		if err != nil {
			continue
		}
		all2.Carts = append(all2.Carts, cartInfo)
	}
	return nil
}

func New(service service.ICartService) *Cart {
	return &Cart{service}
}
