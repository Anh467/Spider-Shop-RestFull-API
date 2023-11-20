package entity

type UserCreate struct {
	Account  string `json:"account" gorm:"column:Account"`
	Password string `json:"password" gorm:"column:Password"`
	Name     string `json:"name" gorm:"column:Name"`
	Gender   string `json:"gender" gorm:"column:Gender"`
	Mail     string `json:"mail" gorm:"column:Mail"`
	Birthday string `json:"birthday" gorm:"column:Birthday"`
}

func (UserCreate) TableModel() string {
	return USER_TABLE
}
