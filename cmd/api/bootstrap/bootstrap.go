package bootstrap

import (
	"fmt"
	"log"

	"api-billing/insfraestructure/handler"
)

func Run() error {
	config := newConfiguration("./configuration.json")
	api := newEcho()

	db, err := newDBConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	handler.InitRoutes(api, db)

	port := fmt.Sprintf(":%d", config.ServerPort)

	return api.Start(port)
}
