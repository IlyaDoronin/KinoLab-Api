package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var cfg conf

var (
	//Storage - объедок хранящий свойства конфига БД
	Storage = &cfg.Storage
	//Server - объедок хранящий свойства конфига сервера
	Server = &cfg.Server
)

//Load - Загружает json файл и десиреализует его
func Load(path string) {
	//Читает файл и проверяет на наличие ошибок
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	//Десиреализует его
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatal(err)
	}
}
