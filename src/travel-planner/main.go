package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-planner/backend"
	"travel-planner/handler"
	"travel-planner/util"
)

func main() {
	fmt.Println("started-service")
	config, err := util.LoadApplicationConfig("conf", "deploy.yml")
	if err != nil {
		panic(err)
	}

	backend.InitMySQLBackend(config.MySQLConfig)
	res, err := backend.DB.exampleQueryFunc("temp")
	log.Fatal(http.ListenAndServe(":8080", handler.InitRouter(config.TokenConfig)))
}
