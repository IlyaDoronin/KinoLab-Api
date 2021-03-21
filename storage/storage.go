package storage

import (
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx"
	"github.com/yeahyeahcore/KinoLab-Api/conf"
)

var (
	conn *pgx.Conn
)

//Init - Инициализация драйвера бд
func Init() {
	fmt.Println("\nInit storage...")

	fmt.Println(GetConfig())

	if !IsConnect() {
		fmt.Println("Подключение было отключено в файлах конфигурации")

		//Ждём 5 сек
		duration := time.Duration(5) * time.Second
		time.Sleep(duration)
		return
	}

	connCfg, err := pgx.ParseURI(GetConfig())
	if err != nil {
		log.Fatal(err)
	}

	conn, err = pgx.Connect(connCfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done.")

}

//GetConfig Возвращает строку конфигурации
//подключения к бд исходя из данных
func GetConfig() string {

	if conf.Storage.Mod == "" {
		return fmt.Sprintf("%s://%s:%s@%s:%s/%s",
			conf.Storage.Driver, conf.Storage.Username,
			conf.Storage.Password, conf.Storage.Host,
			conf.Server.Port, conf.Storage.DBName,
		)
	}

	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?%s",
		conf.Storage.Driver, conf.Storage.Username,
		conf.Storage.Password, conf.Storage.Host,
		conf.Storage.Port, conf.Storage.DBName,
		conf.Storage.Mod,
	)

}

//IsConnect исходя из конфига проверяет
//можно ли подключаться к бд
func IsConnect() bool {
	if conf.Storage.Status == "on" {
		return true
	}
	return false
}
