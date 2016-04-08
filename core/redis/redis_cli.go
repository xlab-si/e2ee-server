package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
	"fmt"
)

type RedisCli struct {
	conn redis.Conn
}

var instanceRedisCli *RedisCli = nil

func Connect() (conn *RedisCli) {
	if instanceRedisCli == nil {
		instanceRedisCli = new(RedisCli)
		
		viper.SetConfigName("config") 
		viper.AddConfigPath("$GOPATH/src/github.com/mancabizjak/e2ee-server/")
	 
		conf_err := viper.ReadInConfig()
		if conf_err != nil {
			fmt.Println(conf_err)
		}
		
		// Read values from config file
		var port = viper.GetInt("redis.port")
		var network_mode = viper.GetString("redis.mode")
		var pw = viper.GetString("redis.password")
		
		var err error
		instanceRedisCli.conn, err = redis.Dial(network_mode, fmt.Sprintf(":%d", port))

		if err != nil {
			panic(err)
		}

		if _, err := instanceRedisCli.conn.Do("AUTH", pw); err != nil {
			instanceRedisCli.conn.Close()
			panic(err)
		}
	}

	return instanceRedisCli
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err == nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	return redisCli.conn.Do("GET", key)
}
