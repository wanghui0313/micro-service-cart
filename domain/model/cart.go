package model

type Cart struct {
	Id        int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ProductId int64 `gorm:"not_null" json:"productId"`
	Num       int64 `gorm:"not_null" json:"num"`
	SizeId    int64 `gorm:"not_null" json:"sizeId"`
	UserId    int64 `gorm:"not_null" json:"userId"`
}
