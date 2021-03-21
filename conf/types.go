package conf

//Конструктор пакета
//omitempy если ноль будет то ничего не запишет
type conf struct {
	Storage storage `json:"storage,omitempty"`
	Server  server  `json:"server,omitempty"`
}

//структура конфига бд
type storage struct {
	Driver   string `json:"driver,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	DBName   string `json:"dbname,omitempty"`
	Mod      string `json:"modification,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Status   string `json:"status,omitempty"`
}

//структура конфига сервера
type server struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}
