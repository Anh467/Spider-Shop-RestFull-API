package common

type AppConext struct {
	AllowHeaders    string    `json:"allowheaders"`
	AllowAllOrigins bool      `json:"allowallorigins"`
	Port            string    `json:"port"`
	SuperSecretKey  string    `json:"supersecretkey"`
	MongoDB         MongoDB   `json:"mongodb"`
	SqlServer       SqlServer `json:"sqlserver"`
	MySQL           MySQL     `json:"mysql"`
}

type MongoDB struct {
	Url string `json:"url"`
	DB  string `json:"db"`
}

type SqlServer struct {
	Host    string `json:"host"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	DB      string `json:"db"`
	Options string `json:"options"`
	Dialect string `json:"dialect"`
}

type MySQL struct {
	Host    string `json:"host"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	DB      string `json:"db"`
	Options string `json:"options"`
	Dialect string `json:"dialect"`
}
