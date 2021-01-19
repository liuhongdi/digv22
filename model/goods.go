package model

type Goods struct {
	GoodsId	int64  `gorm:"column:goodsId",json:"goodsid"` // 自增
	GoodsName string  `gorm:"column:goodsName",json:"goodsname"` //
	Subject string `gorm:"column:subject",json:"subject"`
	Price string `gorm:"column:price",json:"price"`
	Stock string `gorm:"column:stock",json:"stock"`
}

func (Goods) TableName() string {
	return "goods"
}


