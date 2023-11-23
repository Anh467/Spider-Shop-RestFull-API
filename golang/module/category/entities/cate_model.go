package entities

type CateModel struct {
	CateID       int    `json:"cateid" gorm:"column:CateID;primaryKey"`
	Name         string `json:"name" gorm:"column:Name"`
	Desc         string `json:"desc" gorm:"column:Desc;default:''"`
	Image        string `json:"image" gorm:"column:Image;default:''"`
	Status       string `json:"status" gorm:"column:Status;default:Normal"`
	Revenue      int    `json:"revenue" gorm:"column:Revenue"`
	NumWareHouse int    `json:"numwarehouse" gorm:"column:NumWareHouse"`
	NumOrder     int    `json:"numorder" gorm:"column:NumOrder"`
}

func (CateModel) TableName() string {
	return CATE_TABLE
}
