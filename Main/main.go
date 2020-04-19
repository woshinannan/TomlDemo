package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Redis ConfigRedis
}

type ConfigRedis struct {
	Host     string
	Pass     string
	Index    int
	Poolsize int
}

var ConfigValue = new (Config)

func main() {
	_, err := toml.DecodeFile("../config/config.toml", ConfigValue)
	if err != nil {
		panic(err)
	}
	fmt.Printf("host=%s, pass=%s, index=%d, poolsize=%d\n", ConfigValue.Redis.Host, ConfigValue.Redis.Pass, ConfigValue.Redis.Index, ConfigValue.Redis.Poolsize)

	fmt.Printf("Redis=%+v\n", *ConfigValue)


}
