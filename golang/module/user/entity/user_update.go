package entity

type UserUpdate struct {
	Account  string `json:"account" gorm:"column:Account"`
	Password string `json:"password" gorm:"column:Password"`
	Name     string `json:"name" gorm:"column:Name"`
	Image    string `json:"image" gorm:"column:Image"`
	Gender   string `json:"gender" gorm:"column:Gender"`
	Phone    string `json:"phone" gorm:"column:Phone"`
	Mail     string `json:"mail" gorm:"column:Mail"`
	Birthday string `json:"birthday" gorm:"column:Birthday"`
}

func (UserUpdate) TableModel() string {
	return USER_TABLE
}
