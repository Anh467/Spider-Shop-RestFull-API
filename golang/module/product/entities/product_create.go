package entities

type ProductCreate struct {
	CateID    int    `json:"cateid" gorm:"column:CateID"`
	Name      string `json:"name" gorm:"column:Name"`
	Desc      string `json:"desc" gorm:"column:Desc"`
	Image     string `json:"image" gorm:"column:Image"`
	Status    string `json:"status" gorm:"column:Status"`
	CreatedAt string `json:"createdat" gorm:"column:createdAt"`
	UpdatedAt string `json:"updatedat" gorm:"column:updatedAt"`
}

func (ProductCreate) TableName() string {
	return PRODUCT_TABLE
}
