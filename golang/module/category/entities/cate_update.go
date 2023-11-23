package entities

type CateUpdate struct {
	Name   string `json:"name" gorm:"column:Name"`
	Desc   string `json:"desc" gorm:"column:Desc;default:''"`
	Image  string `json:"image" gorm:"column:Image;default:''"`
	Status string `json:"status" gorm:"column:Status;default:Normal"`
}

func (CateUpdate) TableName() string {
	return CATE_TABLE
}
