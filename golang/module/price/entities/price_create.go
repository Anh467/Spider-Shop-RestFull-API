package entities

type PriceCreate struct {
	PriceID   int     `json:"priceid" gorm:"column:PriceID;primaryKey;autoIncrement"`
	ProductID int     `json:"ProductID" gorm:"column:ProductID"`
	Name      string  `json:"name" gorm:"column:Name"`
	Price     float64 `json:"price" gorm:"column:Price"`
	Desc      string  `json:"desc" gorm:"column:Desc"`
	Image     string  `json:"image" gorm:"column:Image"`
	Status    string  `json:"status" gorm:"column:Status;default:Normal"`
}
