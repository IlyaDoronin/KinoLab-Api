package main

import (
	"github.com/yeahyeahcore/KinoLab-Api/conf"
	"github.com/yeahyeahcore/KinoLab-Api/server"
	"github.com/yeahyeahcore/KinoLab-Api/storage"
)

func main() {
	conf.Load("conf/config.json")
	storage.Init()
	server.Start()
}
