package entities

type CateCreate struct {
	Name  string `json:"name" gorm:"column:Name"`
	Desc  string `json:"desc" gorm:"column:Desc;default:''"`
	Image string `json:"image" gorm:"column:Image;default:''"`
}

func (CateCreate) TableName() string {
	return CATE_TABLE
}
