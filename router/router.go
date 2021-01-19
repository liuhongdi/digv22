package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv22/controller"
	"github.com/liuhongdi/digv22/global"
	"log"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)

	/*
	// 路径映射
	goodsc:=controller.NewGoodsController()
	router.GET("/goods/getone/:id", goodsc.GetOne);
	//router.GET("/article/list", articlec.GetList);

	orderc:=controller.NewOrderController()
	router.GET("/order/addone", orderc.AddOne);
	router.GET("/order/addonetx", orderc.AddOneTx);
	*/

	filec:=controller.NewFileController()
	router.GET("/file/getone", filec.GetOne);
	router.GET("/file/downone", filec.DownOne);

	return router
}

func HandleNotFound(c *gin.Context) {
	global.NewResult(c).Error(404,"资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(c).Error(500,"服务器内部错误")
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}