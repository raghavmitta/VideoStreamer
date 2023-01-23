package main

import (
	"VideoStreamer/config"
	"VideoStreamer/db"
	"VideoStreamer/handler"
	"VideoStreamer/service"
)

func main() {
	config.LoadConfig("./config.yml")
	db.Initialize()
	service.StartService(config.GetConfig().Ticker.Time)
	handler.Handle()

}
