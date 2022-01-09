package bootstrap

import (
	"fmt"
	"log"

	"api-billing/insfraestructure/handler"
	"api-billing/insfraestructure/handler/response"
)

func Run() error {
	config := newConfiguration("./configuration.json")
	api := newEcho(config, response.HTTPErrorHandler)

	db, err := newDBConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(api, db)

	port := fmt.Sprintf(":%d", config.ServerPort)

	return api.Start(port)
}
