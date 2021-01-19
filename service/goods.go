package service

import (
	"github.com/liuhongdi/digv22/dao"
	"github.com/liuhongdi/digv22/model"
)

//得到多件商品返回
func GetGoodsList() ([]*model.Goods,error) {
	goods, err := dao.SelectAllGoods()
    return goods,err
}
