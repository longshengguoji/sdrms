package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"sdrms/enums"
	"sdrms/models"
	"strconv"
	"strings"
	"sdrms/utils"
	"time"
)

type FabricController struct {
	BaseController
}

func (c *FabricController) List()  {
	//直接反序化获取json格式的requestbody里的值（要求配置文件里 copyrequestbody=true）
	var params models.FabricListQueryParam
	json.Unmarshal(c.Ctx.Input.RequestBody, &params)
	//获取数据列表和总数
	data, total := models.List(&params)

	//定义返回的数据结构
	//定义返回的数据结构
	result := make(map[string]interface{})
	//resData := make(map[string]interface{})


	result["total"] = total
	result["rows"] = data
	c.Data["json"] = result
	c.ServeJSON()
	}

//Edit 添加、编辑角色界面
func (c *FabricController) Edit() {
	if c.Ctx.Request.Method == "POST" {
		c.Save()
	}
	Id,_:= c.GetInt(":id",0)
	m := models.Fabric{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("数据无效，请刷新后重试")
		}
	}
	c.Data["m"] = m
	c.setTpl("role/edit.html", "shared/layout_pullbox.html")
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footerjs"] = "role/edit_footerjs.html"
}

//Save 添加、编辑页面 保存
func (c *FabricController) Save() {
	var err error
	m := models.Fabric{}
	//获取form里的值
	if err = c.ParseForm(&m); err != nil {
		c.jsonResult(enums.JRCodeFailed, "获取数据失败", m.Id)
	}
	o := orm.NewOrm()
	if m.Id == 0 {
		qrcodeStr,id,err := utils.GetQRCode()
		if err == nil{
			m.BarCode = qrcodeStr
			m.Id = id
			if _, err = o.Insert(&m); err == nil {
				c.jsonResult(enums.JRCodeSucc, "添加成功", m.Id)
			} else {
				c.jsonResult(enums.JRCodeFailed, "添加失败", m.Id)
			}
		}else {
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}

	} else {
		if _, err = o.Update(&m); err == nil {
			c.jsonResult(enums.JRCodeSucc, "编辑成功", m.Id)
		} else {
			c.jsonResult(enums.JRCodeFailed, "编辑失败", m.Id)
		}
	}

}

//Index 角色管理首页
func (c *FabricController) Index() {
	//是否显示更多查询条件的按钮
	c.Data["showMoreQuery"] = true
	//将页面左边菜单的某项激活
	c.Data["activeSidebarUrl"] = c.URLFor(c.controllerName + "." + c.actionName)
	c.setTpl()
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "role/index_headcssjs.html"
	c.LayoutSections["footerjs"] = "role/index_footerjs.html"
	//页面里按钮权限控制
	c.Data["canEdit"] = c.checkActionAuthor("RoleController", "Edit")
	c.Data["canDelete"] = c.checkActionAuthor("RoleController", "Delete")
	c.Data["canAllocate"] = c.checkActionAuthor("RoleController", "Allocate")
}

//Delete 批量删除
func (c *FabricController) Delete() {
	strs := c.GetString("ids")
	ids := make([]int, 0, len(strs))
	for _, str := range strings.Split(strs, ",") {
		if id, err := strconv.Atoi(str); err == nil {
			ids = append(ids, id)
		}
	}
	if num, err := models.FabricBatchDelete(ids); err == nil {
		c.jsonResult(enums.JRCodeSucc, fmt.Sprintf("成功删除 %d 项", num), 0)
	} else {
		c.jsonResult(enums.JRCodeFailed, "删除失败", 0)
	}
}

func (c *FabricController) Print(){
	Id,_:= c.GetInt("id",0)

	//查询对于fabric
	m := models.Fabric{Id: Id}
	if Id > 0 {
		o := orm.NewOrm()
		err := o.Read(&m)
		if err != nil {
			c.pageError("打印出错，请刷新后重试")
		}
	}

	timeFormat := "2006-01-02"
	dateLabel := time.Now().Format(timeFormat)

	tableParam := utils.TablePic{
		Col1:"面料编号NO",
		Val1:m.SampleNumber,
		Col2:"名称Designation",
		Val2:m.ChineseName,
		Col3:"款号",
		Val3:m.StyleNumber,
		Col4:"颜色Color",
		Val4:m.ColorPattern,
		Col5:"克重Weight",
		Val5:m.Weight,
		Col6:"规格Spec",
		Val6:m.Specification,
		Col7:"成份Comp",
		Val7:m.Element,
		Col8:"日期Date",
		Val8:dateLabel,
		Col9:"备注",
		Val9:m.Label,
		DrawText:m.BarCode,
	}
	//qrPicPath := "static/printImg/textQrCode.jpg"
	//tablePicPath := "static/printImg/tablePic.jpg"
	resPicPath := "static/printImg/output.jpg"
	//utils.GenerateQRCode(tableParam.DrawText , dateLabel,qrPicPath)
	utils.GenerateQRAndTablePicture(tableParam,resPicPath)

	picPath := make(map[string]string)
	result := make(map[string]interface{})
	//resData := make(map[string]interface{})
	picPath["qrPicPath"] = resPicPath
	picPath["tablePicPath"] = resPicPath
	result["data"] = picPath
	result["success"] = 0
	c.Data["json"] = result
	c.ServeJSON()
	}
