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
	qrCode,_ = barcode.Scale(qrCode,140,50)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",8,image.Black,140)
	//textBrush, err := NewTextBrush("MicrosoftYaHei.ttf",7,image.Black,140)
	if err != nil {
		fmt.Println(err)
	}
	textBrush.FontSize = 8

	backgroundImg, err := imaging.Open("utils/background.jpg")
	//backgroundImg, err := imaging.Open("background.jpg")
	if err != nil {
		fmt.Println(err)
	}

	dateText := "日期 Date: "+ dateLabel

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(8,12)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(140,12)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(8,118)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(140,118)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(8,223)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(140,223)),qrCode, image.ZP, draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(40,70),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(180,70),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(40,175),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(180,175),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(40,280),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(180,280),encFNC1)

	textBrush.DrawFontOnRGBA(m,image.Pt(10,87),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(152,87),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(10,192),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(152,192),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(10,299),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(152,299),dateText)


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

	qrCode,_ = barcode.Scale(qrCode,140,34)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",8,image.Black,300)
	//textBrush, err := NewTextBrush("MicrosoftYaHei.ttf",8,image.Black,140)
	if err != nil {
		fmt.Print(err)
	}
	textBrush.FontSize = 8

	backgroundImg, err := imaging.Open("utils/tablebg.jpg")
	//backgroundImg, err := imaging.Open("tablebg.jpg")
	if err != nil {
		fmt.Print(err)
	}

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(25,23),tablePic.Col1)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,23),tablePic.Val1)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,50),tablePic.Col2)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,50),tablePic.Val2)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,77),tablePic.Col3)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,77),tablePic.Val3)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,104),tablePic.Col4)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,104),tablePic.Val4)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,131),tablePic.Col5)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,131),tablePic.Val5)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,158),tablePic.Col6)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,158),tablePic.Val6)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,185),tablePic.Col7)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,185),tablePic.Val7)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,212),tablePic.Col8)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,212),tablePic.Val8)
	textBrush.DrawFontOnRGBA(m,image.Pt(25,239),tablePic.Col9)
	textBrush.DrawFontOnRGBA(m,image.Pt(140,239),tablePic.Val9)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(60,270)),qrCode,image.ZP,draw.Src)
	textBrush.DrawFontOnRGBA(m,image.Pt(100,307),tablePic.DrawText)

	imaging.Save(m,picPath)
}
