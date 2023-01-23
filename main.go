package main

import (
	"VideoStreamer/handler"
	"VideoStreamer/service"
)

func main() {
	service.StartService(60)
	handler.Handle()

}
