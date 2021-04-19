package main

import (
	"fmt"
	"sync"
)

// 单例模式

var config *Config
var once sync.Once

type Config struct {
	Port int
}

func GetConfig() *Config {
	once.Do(func() {
		config = &Config{Port: 8080}
	})

	return config
}

func main()  {
	c1 := GetConfig()
	c2 := GetConfig()

	fmt.Println(c1 == c2)
}