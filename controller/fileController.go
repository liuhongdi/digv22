package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv22/service"
	"github.com/signintech/gopdf"
	"log"
	"strconv"
)

type FileController struct{}
func NewFileController() FileController {
	return FileController{}
}
//生成一个pdf文件
func (g FileController) GetOne(c *gin.Context) {
    //创建pdf
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })
	pdf.AddPage()
	//设置font
	err := pdf.AddTTFFont("wts11", "/data/temp/PingFang_Regular.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetTextColor(0x00,0x00,0xff)
	//添加text内容
	pdf.SetX(288)
	pdf.SetY(20)
	pdf.Cell(nil, "您好")

	//添加图片
	pdf.Image("/data/liuhongdi/digv22/static/1200.jpeg", 47, 50, &gopdf.Rect{W: 500, H: 300}) //print image

	//添加表格的标题
	pdf.SetX(47)
	pdf.SetY(370)

	alignOption := gopdf.CellOption{Align: gopdf.Center | gopdf.Middle,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,};

	pdf.CellWithOption(&gopdf.Rect{
		W: 500,
		H: 50,
	}, "旗舰店12月入库商品列表", alignOption)

    //设置表格表头的字体
	err = pdf.SetFont("wts11", "", 10)
	pdf.SetTextColor(0xa8,0xa8,0xa8)
	if err != nil {
		log.Print(err.Error())
		return
	}

	//生成表格的表头
	pdf.SetX(47)
	pdf.SetY(420)
	pdf.CellWithOption(&gopdf.Rect{
		W: 50,
		H: 40,
	}, "商品ID", alignOption)

	pdf.SetX(97)
	pdf.SetY(420)

	pdf.CellWithOption(&gopdf.Rect{
		W: 400,
		H: 40,
	}, "商品名称", alignOption)

	pdf.SetX(497)
	pdf.SetY(420)

	pdf.CellWithOption(&gopdf.Rect{
		W: 50,
		H: 40,
	}, "商品库存", alignOption)


	//生成表格的内容
	pdf.SetTextColor(0x00,0x00,0x00)
	alignLeft := gopdf.CellOption{Align: gopdf.Left | gopdf.Middle,
		Border: gopdf.Left | gopdf.Right | gopdf.Bottom | gopdf.Top,};
	//查询数据库，用表格显示数据
	//得到商品列表
	goods,err := service.GetGoodsList()
	//遍历商品
	for i,v := range goods {
		   curY := 460+i*40

		   pdf.SetX(47)
		   pdf.SetY(float64(curY))

		   pdf.CellWithOption(&gopdf.Rect{
			   W: 50,
			   H: 40,
		   }, " "+strconv.FormatInt(v.GoodsId,10), alignLeft)

		   pdf.SetX(97)
		   pdf.SetY(float64(curY))

		   pdf.CellWithOption(&gopdf.Rect{
			   W: 400,
			   H: 40,
		   }, " "+v.GoodsName, alignLeft)

		   pdf.SetX(497)
		   pdf.SetY(float64(curY))

		   pdf.CellWithOption(&gopdf.Rect{
			   W: 50,
			   H: 40,
		   }, " "+string(v.Stock), alignLeft)
	}
	//save to file
	pdf.WritePdf("/data/temp/hello.pdf")
	return
}

//生成一个pdf文件并下载
func (g FileController) DownOne(c *gin.Context) {

	filepath:="/data/temp/hello.pdf"
    filename:="hello.pdf"
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))//fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "text/pdf")
	c.File(filepath)

	return
}