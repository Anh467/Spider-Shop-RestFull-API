package entities

type UserUpdate struct {
	Password string `json:"password" gorm:"column:Password"`
	Name     string `json:"name" gorm:"column:Name"`
	Image    string `json:"image" gorm:"column:Image"`
	Gender   string `json:"gender" gorm:"column:Gender"`
	Phone    string `json:"phone" gorm:"column:Phone"`
	Mail     string `json:"mail" gorm:"column:Mail"`
	Birth    string `json:"birth" gorm:"column:Birth"`
}

func (UserUpdate) TableModel() string {
	return USER_TABLE
}
