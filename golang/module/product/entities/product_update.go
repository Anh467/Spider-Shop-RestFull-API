package entities

type ProductUpdate struct {
	CateID int    `json:"cateid" gorm:"column:CateID"`
	Name   string `json:"name" gorm:"column:Name"`
	Desc   string `json:"desc" gorm:"column:Desc"`
	Image  string `json:"image" gorm:"column:Image"`
	Status string `json:"status" gorm:"column:Status;default:Normal"`
}

func (ProductUpdate) TableName() string {
	return PRODUCT_TABLE
}
