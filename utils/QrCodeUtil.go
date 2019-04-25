package utils

import (
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
)


func GenerateQRCode(fnCode string, dateLabel string,resPicPath string){
	encFNC1 := fnCode
	qrCode,_ := code128.EncodeWithoutChecksum(encFNC1)
	qrCode,_ = barcode.Scale(qrCode,450,260)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",14,image.Black,300)
	if err != nil {
		fmt.Println(err)
	}
	textBrush.FontSize = 14

	backgroundImg, err := imaging.Open("utils/background.jpg")
	if err != nil {
		fmt.Println(err)
	}

	dateText := "日期 Date: "+ dateLabel

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(30,40)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(525,40)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(30,415)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(525,415)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(30,785)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(525,785)),qrCode, image.ZP, draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(205,310),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(700,310),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(205,685),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(700,685),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(205,1055),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(700,1055),encFNC1)

	textBrush.DrawFontOnRGBA(m,image.Pt(155,350),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(650,350),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(155,725),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(650,725),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(155,1095),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(650,1095),dateText)


	imaging.Save(m,resPicPath)
}

type TablePic struct {
	Col1 string
	Val1 string
	Col2 string
	Val2 string
	Col3 string
	Val3 string
	Col4 string
	Val4 string
	Col5 string
	Val5 string
	Col6 string
	Val6 string
	Col7 string
	Val7 string
	Col8 string
	Val8 string
	Col9 string
	Val9 string
	DrawText string
}

func GenerateTablePicture(tablePic TablePic,picPath string){
	encFNC1 := tablePic.DrawText
	qrCode, _ := code128.EncodeWithoutChecksum(encFNC1)

	qrCode,_ = barcode.Scale(qrCode,500,100)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",14,image.Black,300)
	if err != nil {
		fmt.Print(err)
	}

	backgroundImg, err := imaging.Open("utils/tablebg.jpg")
	if err != nil {
		fmt.Print(err)
	}

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(80,100),tablePic.Col1)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,100),tablePic.Val1)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,200),tablePic.Col2)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,200),tablePic.Val2)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,290),tablePic.Col3)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,290),tablePic.Val3)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,375),tablePic.Col4)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,375),tablePic.Val4)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,470),tablePic.Col5)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,470),tablePic.Val5)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,560),tablePic.Col6)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,560),tablePic.Val6)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,650),tablePic.Col7)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,650),tablePic.Val7)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,740),tablePic.Col8)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,740),tablePic.Val8)
	textBrush.DrawFontOnRGBA(m,image.Pt(80,830),tablePic.Col9)
	textBrush.DrawFontOnRGBA(m,image.Pt(480,830),tablePic.Val9)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(220,920)),qrCode,image.ZP,draw.Src)
	textBrush.DrawFontOnRGBA(m,image.Pt(400,1070),tablePic.DrawText)

	imaging.Save(m,picPath)
}
