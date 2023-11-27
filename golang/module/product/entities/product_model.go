package entities

import entities_category "SpiderShop-Restfull-API/module/category/entities"

type ProductModel struct {
	ProductID int                              `json:"productid" gorm:"column:ProductID;primaryKey"`
	CateID    int                              `json:"cateid" gorm:"column:CateID"`
	Cate      entities_category.CateGetProduct `json:"cate" gorm:"foreignKey:CateID"`
	Name      string                           `json:"name" gorm:"column:Name"`
	Desc      string                           `json:"desc" gorm:"column:Desc"`
	Image     string                           `json:"image" gorm:"column:Image"`
	Status    string                           `json:"status" gorm:"column:Status;default:Normal"`
	CreatedAt string                           `json:"createdat" gorm:"column:createdAt"`
	UpdatedAt string                           `json:"updatedat" gorm:"column:updatedAt"`
}

func (ProductModel) TableName() string {
	return PRODUCT_TABLE
}
