package entity

type UserModel struct {
	UserID      int    `json:"user_id" gorm:"column:UserID;primaryKey"`
	Account     string `json:"account" gorm:"column:Account"`
	Password    string `json:"password" gorm:"column:Password"`
	Name        string `json:"name" gorm:"column:Name"`
	Image       string `json:"image" gorm:"column:Image"`
	Gender      string `json:"gender" gorm:"column:Gender"`
	Phone       string `json:"phone" gorm:"column:Phone"`
	Mail        string `json:"mail" gorm:"column:Mail"`
	Birth       string `json:"birth" gorm:"column:Birth"`
	Role        string `json:"role" gorm:"column:Role"`
	Expenditure string `json:"expenditure" gorm:"column:Expenditure"`
	NumOrder    int    `json:"numorder" gorm:"column:NumOrder"`
	CreatedAt   string `json:"created_at" gorm:"column:CreatedAt"`
	UpdatedAt   string `json:"updated_at" gorm:"column:UpdatedAt"`
}

func (UserModel) TableModel() string {
	return USER_TABLE
}
