package entities

type UserCredential struct {
	Account  string `json:"account" gorm:"column:Account"`
	Password string `json:"password" gorm:"column:Password"`
}
