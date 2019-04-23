package models

import (
	"github.com/astaxie/beego/orm"
)

// FabricListQueryParam 用于查询的类
type FabricListQueryParam struct {
	BaseQueryParam
	ProductCategoryLike string `json:"productCategoryLike"` //模糊查询
	BarCodeIn string  `json:"barCodeIn"` //模糊查询
}

type Fabric struct {
	Id int `json:"id";orm:"column(id);auto"`
	Label string `json:"label"`
	Catalogue string `json:"catalogue"`
	SubCatalogue string `json:"subCatalogue";orm:"column(sub_catalogue)"`
	BarCode string `json:"barCode";orm:"column(bar_code)"`
	ChineseName string `json:"chineseName"`
	SampleNumber string `json:"sampleNumber"`
	Description string `json:"description"`
	ColorPattern string `json:"colorPattern"`
	StyleNumber string `json:"styleNumber"`
	Specification string `json:"specification"`
	GreigeSupplier string `json:"greigeSupplier"`
	ProductCategory string `json:"productCategory"`
	Supplier string `json:"supplier"`
	UnitPrice string `json:"unitPrice";orm:"column(unit_price)"`
	Element string `json:"element"`
	Weight string `json:"weight"`
}

func (m *Fabric) TableName() string {
	return FabricListTBName()
}

func  List(params *FabricListQueryParam) ([]*Fabric,int64)  {
	query := orm.NewOrm().QueryTable(FabricListTBName())

	fabricList := make([]*Fabric,0)

	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	if len(params.ProductCategoryLike) > 0 {
		query = query.Filter("product_category", params.ProductCategoryLike)
	}
	if len(params.BarCodeIn) > 0 {
		query = query.Filter("bar_code__contains", params.BarCodeIn)
	}
	total, _ := query.Count()
	query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&fabricList)

	return fabricList, total
	}

// FabricBatchDelete 批量删除
func FabricBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(FabricListTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}