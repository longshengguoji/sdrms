package utils

import (
	"testing"
)

//func TestGenerateQRCode(t *testing.T){
//	fnCode := "m171126391"
//	timeFormat := "2006-01-02"
//	dateLabel := time.Now().Format(timeFormat)
//	picPath := "qrcode.jpg"
//
//	GenerateQRCode(fnCode , dateLabel,picPath)
//}

func TestGenerateTablePicture(t *testing.T){
	tableParam := TablePic{
		Col1:"面料编号NO",
		Val1:"SC-P-12-000509",
		Col2:"名称Designation",
		Val2:"FTY-3038",
		Col3:"门幅Width",
		Val3:"140",
		Col4:"颜色Color",
		Val4:"黑色",
		Col5:"克重Weight",
		Val5:"",
		Col6:"规格Spec",
		Val6:"",
		Col7:"成份Comp",
		Val7:"100%P",
		Col8:"日期Date",
		Val8:"2019-08-18",
		Col9:"备注",
		Val9:"TED",
		DrawText:"M171126391",
	}
	picPath := "tablepic.jpg"

	GenerateQRAndTablePicture(tableParam,picPath)
}
