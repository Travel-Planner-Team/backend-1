package main

import (
	"fmt"
	"log"
	"net/http"
	"travel-planner/util"
	"travel-planner/backend"
	"travel-planner/handler"
)

func main() {
	fmt.Println("started-service")
	config, err := util.LoadApplicationConfig("conf", "deploy.yml")
    if err != nil {
        panic(err)
    }

    // backend.InitMySQLBackend(config.ElasticsearchConfig)

    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter(config.TokenConfig)))
}
