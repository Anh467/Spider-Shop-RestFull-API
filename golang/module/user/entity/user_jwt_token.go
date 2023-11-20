package entity

type UserJWTModel struct {
	UserID   int    `json:"user_id" gorm:"column:UserID;primaryKey"`
	Account  string `json:"account" gorm:"column:Account"`
	Name     string `json:"name" gorm:"column:Name"`
	Gender   string `json:"gender" gorm:"column:Gender"`
	Mail     string `json:"mail" gorm:"column:Mail"`
	Birthday string `json:"birthday" gorm:"column:Birthday"`
	Token    string `json:"token"`
}

func (UserJWTModel) TableModel() string {
	return USER_TABLE
}