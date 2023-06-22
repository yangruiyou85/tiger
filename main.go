// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"github.com/yangruiyou85/tiger/backup/api"
)

func main() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	http.ListenAndServe(":8080", api.InitRouter())
}
