package utils

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/disintegration/imaging"
	"image"
	"image/draw"
	"reflect"
	"strconv"
	"time"
)

type QrCode struct{
	Year int `json:"year"`
	Code int `json:"code"`
	Id int `json:"id"`
}

func GenerateQRCode(fnCode string, dateLabel string,resPicPath string){
	encFNC1 := fnCode
	qrCode,_ := code128.EncodeWithoutChecksum(encFNC1)
	qrCode,_ = barcode.Scale(qrCode,550,300)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",35,image.Black,140)
	//textBrush, err := NewTextBrush("MicrosoftYaHei.ttf",35,image.Black,700)
	if err != nil {
		fmt.Println(err)
	}
	textBrush.FontSize = 35

	backgroundImg, err := imaging.Open("utils/background.jpg")
	//backgroundImg, err := imaging.Open("background.jpg")
	if err != nil {
		fmt.Println(err)
	}

	dateText := "日期 Date: "+ dateLabel

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(26,105)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(748,105)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(26,710)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(748,710)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(26,1315)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(748,1315)),qrCode, image.ZP, draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(250,530),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(970,530),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(250,1135),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(970,1135),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(250,1740),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(970,1740),encFNC1)

	textBrush.DrawFontOnRGBA(m,image.Pt(130,600),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(850,600),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(130,1205),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(850,1205),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(130,1810),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(850,1810),dateText)


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

	qrCode,_ = barcode.Scale(qrCode,1000,300)

	//textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",8,image.Black,300)
	textBrush, err := NewTextBrush("MicrosoftYaHei.ttf",35,image.Black,500)
	if err != nil {
		fmt.Print(err)
	}
	textBrush.FontSize = 35

	//backgroundImg, err := imaging.Open("utils/background.jpg")
	backgroundImg, err := imaging.Open("background.jpg")
	if err != nil {
		fmt.Print(err)
	}

	m := image.NewRGBA(backgroundImg.Bounds())
	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(1550,130),tablePic.Col1)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,130),tablePic.Val1)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,280),tablePic.Col2)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,280),tablePic.Val2)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,430),tablePic.Col3)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,430),tablePic.Val3)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,580),tablePic.Col4)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,580),tablePic.Val4)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,730),tablePic.Col5)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,730),tablePic.Val5)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,880),tablePic.Col6)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,880),tablePic.Val6)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,1030),tablePic.Col7)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,1030),tablePic.Val7)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,1180),tablePic.Col8)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,1180),tablePic.Val8)
	textBrush.DrawFontOnRGBA(m,image.Pt(1550,1330),tablePic.Col9)
	textBrush.DrawFontOnRGBA(m,image.Pt(2150,1330),tablePic.Val9)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(1700,1450)),qrCode,image.ZP,draw.Src)
	textBrush.DrawFontOnRGBA(m,image.Pt(2050,1800),tablePic.DrawText)

	imaging.Save(m,picPath)
}


func GenerateQRAndTablePicture(tablePic TablePic,picPath string){

	encFNC1 := tablePic.DrawText
	qrCodeInTable, _ := code128.EncodeWithoutChecksum(encFNC1)

	qrCodeInTable,_ = barcode.Scale(qrCodeInTable,800,300)

	textBrush, err := NewTextBrush("utils/MicrosoftYaHei.ttf",35,image.Black,700)
	//textBrush, err := NewTextBrush("MicrosoftYaHei.ttf",30,image.Black,500)
	if err != nil {
		fmt.Print(err)
	}
	textBrush.FontSize = 30

	backgroundImg, err := imaging.Open("utils/background.jpg")
	//backgroundImg, err := imaging.Open("background.jpg")
	if err != nil {
		fmt.Print(err)
	}

	m := image.NewRGBA(backgroundImg.Bounds())

	qrCode,_ := code128.EncodeWithoutChecksum(encFNC1)
	qrCode,_ = barcode.Scale(qrCode,570,300)


	dateText := "日期 Date: "+ tablePic.Val8

	draw.Draw(m,backgroundImg.Bounds(),backgroundImg,image.ZP,draw.Src)

	draw.Draw(m,qrCode.Bounds().Add(image.Pt(40,70)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(650,70)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(40,620)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(650,620)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(40,1170)),qrCode, image.ZP, draw.Src)
	draw.Draw(m,qrCode.Bounds().Add(image.Pt(650,1170)),qrCode, image.ZP, draw.Src)

	textBrush.DrawFontOnRGBA(m,image.Pt(220,420),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(830,420),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(220,970),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(830,970),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(220,1520),encFNC1)
	textBrush.DrawFontOnRGBA(m,image.Pt(830,1520),encFNC1)

	textBrush.DrawFontOnRGBA(m,image.Pt(100,490),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(710,490),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(100,1040),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(710,1040),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(100,1590),dateText)
	textBrush.DrawFontOnRGBA(m,image.Pt(710,1590),dateText)


//table print
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,165),tablePic.Col1)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,165),tablePic.Val1)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,280),tablePic.Col2)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,280),tablePic.Val2)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,395),tablePic.Col3)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,395),tablePic.Val3)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,510),tablePic.Col4)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,510),tablePic.Val4)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,625),tablePic.Col5)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,625),tablePic.Val5)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,740),tablePic.Col6)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,740),tablePic.Val6)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,855),tablePic.Col7)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,855),tablePic.Val7)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,970),tablePic.Col8)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,970),tablePic.Val8)
	textBrush.DrawFontOnRGBA(m,image.Pt(1300,1085),tablePic.Col9)
	textBrush.DrawFontOnRGBA(m,image.Pt(1850,1085),tablePic.Val9)

	draw.Draw(m,qrCodeInTable.Bounds().Add(image.Pt(1450,1200)),qrCodeInTable,image.ZP,draw.Src)
	textBrush.DrawFontOnRGBA(m,image.Pt(1700,1530),tablePic.DrawText)

	imaging.Save(m,picPath)
}

func GetQRCode() (string,int,error) {
	year := time.Now().Year()
	yearNum := year%100
	yearStr := strconv.Itoa(yearNum)

	var qrCode QrCode
	o := orm.NewOrm()
	err := o.Raw("SELECT year, code, id FROM tb_qr_code WHERE year = ?", year).QueryRow(&qrCode)
	newCode := 1111111
	resId := 1
		if err == nil && !reflect.DeepEqual(qrCode, QrCode{}) {
				newCode = qrCode.Code + 1
				resId = qrCode.Id + 1
				qrCode.Code = newCode
				qrCode.Id = resId
			res, err := o.Raw("UPDATE tb_qr_code SET code = ?,id = ? WHERE year = ?", newCode,resId,year).Exec()
			if err == nil {
				num, _ := res.RowsAffected()
				fmt.Println("mysql row affected nums: ", num)
			}
		}else{
			qrCode.Code = newCode
			qrCode.Year = year
			var lastYearQrCode QrCode
			err := o.Raw("SELECT year, code, id FROM tb_qr_code WHERE year = ?", year).QueryRow(&lastYearQrCode)
			if err == nil{
				resId = lastYearQrCode.Id + 1
			}
			qrCode.Id = resId

			res, err := o.Raw("INSERT INTO tb_qr_code (year, code,id) VALUES(?,?,?)", year,newCode,resId).Exec()
			if err == nil {
				num, _ := res.RowsAffected()
				fmt.Println("mysql row affected nums: ", num)
			}

		}
	codeStr := strconv.Itoa(newCode)
	resCode := "M"+ yearStr + codeStr
	return resCode,resId,nil
}