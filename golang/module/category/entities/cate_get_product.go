package entities

type CateGetProduct struct {
	CateID int    `json:"cateid" gorm:"column:CateID;primaryKey"`
	Name   string `json:"name" gorm:"column:Name"`
	Desc   string `json:"desc" gorm:"column:Desc;default:''"`
	Status string `json:"status" gorm:"column:Status;default:Normal"`
}

func (CateGetProduct) TableName() string {
	return CATE_TABLE
}
