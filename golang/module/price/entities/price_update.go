package entities

type PriceUpdate struct {
	ProductID int     `json:"ProductID" gorm:"column:ProductID"`
	Name      string  `json:"name" gorm:"column:Name"`
	Price     float64 `json:"price" gorm:"column:Price"`
	Desc      string  `json:"desc" gorm:"column:Desc"`
	Image     string  `json:"image" gorm:"column:Image"`
	Status    string  `json:"status" gorm:"column:Status;default:Normal"`
}

func (PriceUpdate) TableName() string {
	return Price_TABLE
}
