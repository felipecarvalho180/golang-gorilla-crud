package main

import (
	"api-rest/database"
	"api-rest/models"
	"api-rest/routes"
	"fmt"
)

func main() {
	models.Persons = []models.Person{
		{
			Id:          1,
			Name:        "Paulo Barbosa da Silva",
			Description: "Paulo Barbosa da Silva (Sabará, 25 de janeiro de 1790, 28 de janeiro de 1868) foi mordomo-mor da Casa Imperial do Brasil e deputado pela então província de Minas Gerais. A sua participação na fundação da cidade de Petrópolis foi decisiva quando mobilizou o seu companheiro de arma, o engenheiro major Júlio Frederico Koeler.",
		},
		{
			Id:          2,
			Name:        "Júlio Frederico Koeler",
			Description: "Júlio Frederico Koeler, ou Julius Friedrich Koeler, (Mainz, Grão Ducado de Hessen - Darmstadt, Alemanha, 16 de junho de 1804 – Petrópolis, 1847) foi um militar e engenheiro teuto-brasileiro.",
		},
	}

	database.DBConnect()
	fmt.Println("Starting Server 🚀")
	routes.HandleRequest()
}
