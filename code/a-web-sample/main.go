package main

import (
	"a-web-sample/config"
	"a-web-sample/router"
)

func main() {

	r := router.SetupRouter()

	config.InitConfig()

	port := config.AppConfig.App.Port

	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
