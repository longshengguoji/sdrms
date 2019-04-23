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
	qrCode,_ = barcode.Scale(qrCode,340,75)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",12,image.Black,200)
	if err != nil {
		fmt.Println(err)
	}
	textBrush.FontSize = 12

	backgroundImg, err := imaging.Open("utils/background.jpg")
	if err != nil {
		fmt.Println(err)
	}

	dateText := "日期 Date: "+ dateLabel

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(25,30)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(385,30)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(25,185)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(385,185)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(25,330)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(385,330)),qrCode, image.ZP, draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(150,110),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(500,110),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(150,260),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(500,260),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(150,410),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(500,410),encFNC1)

	textBrush.DrawFontOnRGBA(m,image.Pt(100,135),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(450,135),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(100,285),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(450,285),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(100,435),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(450,435),dateText)


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

	qrCode,_ = barcode.Scale(qrCode,250,50)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",10,image.Black,200)
	if err != nil {
		fmt.Print(err)
	}

	backgroundImg, err := imaging.Open("utils/tablebg.jpg")
	if err != nil {
		fmt.Print(err)
	}

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(60,38),tablePic.Col1)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,38),tablePic.Val1)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,76),tablePic.Col2)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,76),tablePic.Val2)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,114),tablePic.Col3)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,114),tablePic.Val3)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,152),tablePic.Col4)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,152),tablePic.Val4)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,190),tablePic.Col5)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,190),tablePic.Val5)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,228),tablePic.Col6)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,228),tablePic.Val6)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,266),tablePic.Col7)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,266),tablePic.Val7)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,304),tablePic.Col8)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,304),tablePic.Val8)
	textBrush.DrawFontOnRGBA(m,image.Pt(60,342),tablePic.Col9)
	textBrush.DrawFontOnRGBA(m,image.Pt(360,342),tablePic.Val9)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(220,388)),qrCode,image.ZP,draw.Src)
	textBrush.DrawFontOnRGBA(m,image.Pt(300,440),tablePic.DrawText)

	imaging.Save(m,picPath)
}
