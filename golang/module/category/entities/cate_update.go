package entities

type CateUpdate struct {
	Name   string `json:"name" gorm:"column:Name"`
	Desc   string `json:"desc" gorm:"column:Desc"`
	Image  string `json:"image" gorm:"column:Image"`
	Status string `json:"status" gorm:"column:Status"`
}

func (CateUpdate) TableName() string {
	return CATE_TABLE
}
