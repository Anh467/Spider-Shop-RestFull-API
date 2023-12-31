package entities

type UserCreate struct {
	UserID   int    `json:"userid" gorm:"column:UserID;primaryKey"`
	Account  string `json:"account" gorm:"column:Account"`
	Password string `json:"password" gorm:"column:Password"`
	Name     string `json:"name" gorm:"column:Name"`
	Gender   string `json:"gender" gorm:"column:Gender"`
	Mail     string `json:"mail" gorm:"column:Mail"`
	Birth    string `json:"birth" gorm:"column:Birth"`
}

func (UserCreate) TableName() string {
	return USER_TABLE
}
