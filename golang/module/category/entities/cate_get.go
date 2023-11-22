package entities

type CateGet struct {
	CateID       int    `json:"cateid" gorm:"column:CateID;primaryKey"`
	Name         string `json:"name" gorm:"column:Name"`
	Desc         string `json:"desc" gorm:"column:Desc"`
	Image        string `json:"image" gorm:"column:Image"`
	Status       string `json:"status" gorm:"column:Status"`
	Revenue      int    `json:"revenue" gorm:"column:Revenue"`
	NumWareHouse int    `json:"numwarehouse" gorm:"column:NumWareHouse"`
	NumOrder     int    `json:"numorder" gorm:"column:NumOrder"`
}

func (CateGet) TableName() string {
	return CATE_TABLE
}
