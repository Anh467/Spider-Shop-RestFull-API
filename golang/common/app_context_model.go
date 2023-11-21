package common

import "gorm.io/gorm"

type AppConext struct {
	GormDB           *gorm.DB  `json:"gormdb"`
	AllowHeaders     []string  `json:"allowheaders"`
	AllowAllOrigins  bool      `json:"allowallorigins"`
	AllowMethods     []string  `json:"allowmethods"`
	AllowCredentials bool      `json:"allowcredentials"`
	Port             string    `json:"port"`
	SuperSecretKey   string    `json:"secretKey"`
	MongoDB          MongoDB   `json:"mongodb"`
	SqlServer        SqlServer `json:"sqlserver"`
	MySQL            MySQL     `json:"mysql"`
}

type MongoDB struct {
	GormDB *gorm.DB `json:"gormdb"`
	Url    string   `json:"url"`
	DB     string   `json:"db"`
}

type SqlServer struct {
	GormDB  *gorm.DB `json:"gormdb"`
	Host    string   `json:"host"`
	User    string   `json:"user"`
	Pass    string   `json:"pass"`
	DB      string   `json:"db"`
	Options string   `json:"options"`
	Dialect string   `json:"dialect"`
}

type MySQL struct {
	GormDB  *gorm.DB `json:"gormdb"`
	Host    string   `json:"host"`
	User    string   `json:"user"`
	Pass    string   `json:"pass"`
	DB      string   `json:"db"`
	Options string   `json:"options"`
	Dialect string   `json:"dialect"`
}
