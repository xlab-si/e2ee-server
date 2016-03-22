package settings

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
)

type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings Settings = Settings{}
var env = "preproduction"

func Init() {
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
	
	viper.SetConfigName("config") 
	viper.AddConfigPath("$GOPATH/src/github.com/xlab-si/e2ee-server/")
 
	conf_err := viper.ReadInConfig()
	if conf_err != nil {
		fmt.Println(conf_err)
	}

	var env_path = "environments." + env
	settings = Settings{}
	err := viper.UnmarshalKey(env_path, &settings)
	
	if err != nil {
		fmt.Println("Error: Unable to unmarshal key to settings struct")
	}
}

func GetEnvironment() string {
	return env
}

func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

func IsTestEnvironment() bool {
	return env == "tests"
}