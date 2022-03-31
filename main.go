package main

import (
	"github.com/rafaelspirlandelli/api-go-gin/database"
	"github.com/rafaelspirlandelli/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
