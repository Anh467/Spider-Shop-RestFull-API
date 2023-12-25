package entities

import "SpiderShop-Restfull-API/module/product/entities"

type PriceModel struct {
	PriceID      int                 `json:"priceid" gorm:"column:PriceID;primaryKey"`
	ProductID    int                 `json:"ProductID" gorm:"column:ProductID"`
	Product      entities.ProductGet `gorm:"foreignKey:ProductID"`
	Name         string              `json:"name" gorm:"column:Name"`
	Price        float64             `json:"price" gorm:"column:Price"`
	Desc         string              `json:"desc" gorm:"column:Desc"`
	Image        string              `json:"image" gorm:"column:Image"`
	Status       string              `json:"status" gorm:"column:Status;default:Normal"`
	Revenue      float64             `json:"revenue" gorm:"column:Revenue;default:0"`
	NumWareHouse int                 `json:"numwarehouse" gorm:"column:NumWareHouse;default:0"`
	NumOrder     int                 `json:"numorder" gorm:"column:NumOrder;default:0"`
	CreatedAt    string              `json:"createdat" gorm:"column:createdAt"`
	UpdatedAt    string              `json:"updatedat" gorm:"column:updatedAt"`
}

func (PriceModel) TableName() string {
	return Price_TABLE
}
