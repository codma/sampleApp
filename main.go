package main

import (
	"main/api"
	"main/common/setting"
	"main/constant"
	"main/database"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := setting.AppSetting()
	if err != nil {
		panic(err)
	}

	database.Client, err = database.Setup()
	if err != nil {
		panic(err)
	}

}

func main() {
	routersInit := api.InitRouter()
	endPoint := viper.GetString(constant.SERVER_HTTP_PORT)

	routersInit.Listen(endPoint)
}
