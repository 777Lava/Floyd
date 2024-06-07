package main

import (
	"backend"
	"backend/pkg"
)

func main() {
	srv := new(backend.Server)
	services := pkg.NewService()
	handlers := pkg.NewHandler(services)

	srv.Run("8080", handlers.InitRoutes())
}
