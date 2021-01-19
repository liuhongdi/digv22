package dao

import (
	"github.com/liuhongdi/digv22/global"
	"github.com/liuhongdi/digv22/model"
)

func SelectAllGoods() ([]*model.Goods, error) {
	var goods []*model.Goods
	global.DBLink.Find(&goods)
	return goods, nil
}