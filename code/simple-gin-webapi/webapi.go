package main

import (
	"fmt"
	db "webapi/database"
	"webapi/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	db.InitDB()

	user, err := db.GetUserDetail(4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*user)

	addr := viper.GetString("server.addr")

	app := gin.Default()
	app = router.BuildRoute(app)
	app.Run(addr)

}
