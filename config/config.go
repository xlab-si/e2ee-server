package config	

import (
	"github.com/spf13/viper"
	"fmt"
)

func Init() {
	viper.SetConfigName("config") 
	viper.AddConfigPath("$GOPATH/src/github.com/xlab-si/e2ee-server/config/")
 
	conf_err := viper.ReadInConfig()
	if conf_err != nil {
		fmt.Println(conf_err)
	}
}
