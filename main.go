package main

import (
	"go-ms-adapter/config"
	"go-ms-adapter/routes"
)

func main() {
	router := routes.Router()
	s := config.Server{}
	s.Initialize(router)
	s.Run()
}
