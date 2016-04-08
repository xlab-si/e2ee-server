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
	if EnvVariablesSet(env) {
		// Read configuration from environment variables
		fmt.Println("Reading settings from environment variables")
		settings = LoadSettingsFromEnvVars(env)
	} else {
		// Read configuration from config file
		fmt.Println("Environment variables not set, reading settings from config file")
		viper.SetConfigName("config") 
		viper.AddConfigPath("$GOPATH/src/github.com/mancabizjak/e2ee-server/")
	 
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
}

// Checks for the presence of env variables
func EnvVariablesSet(env string) bool {
	viper.SetEnvPrefix("env"); // variables of form ENV_
	viper.AutomaticEnv()
	
	return viper.IsSet(env + "_priv") && viper.IsSet(env + "_pub") && viper.IsSet(env + "_jwt")
}

// Reads configuration settings to global settings struct
func LoadSettingsFromEnvVars(env string) Settings {
	return Settings { PrivateKeyPath: viper.GetString(env + "_priv"),
					 PublicKeyPath: viper.GetString(env + "_pub"),
					 JWTExpirationDelta: viper.GetInt(env + "_jwt") }
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
