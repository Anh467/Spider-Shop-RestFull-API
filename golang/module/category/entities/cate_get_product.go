package entities

type CateGetProduct struct {
	CateID int    `json:"cateid" gorm:"column:CateID;primaryKey"`
	Name   string `json:"name" gorm:"column:Name"`
	Desc   string `json:"desc" gorm:"column:Desc;default:''"`
}

func (CateGetProduct) TableName() string {
	return CATE_TABLE
}
