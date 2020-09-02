package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/hyahm/xredis"
)

type user struct {
	Name  string
	Age   int
	Id    int64
	Role  int
	Token string
}

func main() {
	conf := &redis.Options{
		Network:            "",
		Addr:               "192.168.50.211:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           10,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	client, err := xredis.Conn(conf)
	if err != nil {
		panic(err)
	}
	if _, err := client.NewStr().Set("aa", "bbb", time.Second*10); err != nil {
		panic(err)
	}
	fmt.Println("设置成功")

	value, err := client.NewStr().Get("aa")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

}
