package entities

type UserJWTModel struct {
	UserID  int    `json:"userid" gorm:"column:UserID;primaryKey"`
	Account string `json:"account" gorm:"column:Account"`
	Name    string `json:"name" gorm:"column:Name"`
	Gender  string `json:"gender" gorm:"column:Gender"`
	Mail    string `json:"mail" gorm:"column:Mail"`
	Birth   string `json:"birth" gorm:"column:Birth"`
	Token   string `json:"token"`
}

func (UserJWTModel) TableName() string {
	return USER_TABLE
}
