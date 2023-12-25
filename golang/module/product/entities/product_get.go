package entities

import entitiesCate "SpiderShop-Restfull-API/module/category/entities"

type ProductGet struct {
	ProductID int                         `json:"productid" gorm:"column:ProductID;primaryKey;autoIncrement"`
	CateID    int                         `json:"cateid" gorm:"column:CateID"`
	Cate      entitiesCate.CateGetProduct `gorm:"foreignKey:CateID"`
	Name      string                      `json:"name" gorm:"column:Name"`
	Desc      string                      `json:"desc" gorm:"column:Desc"`
	Image     string                      `json:"image" gorm:"column:Image"`
	Status    string                      `json:"status" gorm:"column:Status;default:Normal"`
	CreatedAt string                      `json:"createdat" gorm:"column:createdAt"`
	UpdatedAt string                      `json:"updatedat" gorm:"column:updatedAt"`
}

func (ProductGet) TableName() string {
	return PRODUCT_TABLE
}
