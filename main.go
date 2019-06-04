package main

import (
	"go/go-adapter-framework/config"
	"go/go-adapter-framework/routes"
)

func main() {
	router := routes.Router()
	s := config.Server{}
	s.Initialize(router)
	s.Run()
}
