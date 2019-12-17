package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/hyahm/xredis"
	"time"
)

func main() {
	conf := &redis.Options{
		Network:            "",
		Addr:               "127.0.0.1:6379",
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
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	client, err  := xredis.Conn(conf)
	if err != nil {
		panic(err)
	}

	if _, err := client.Str().Set("aa", "bbb", time.Second*10); err != nil {
		panic(err)
	}
	fmt.Println("设置成功")
	s, err := client.Str().Get("aa")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	//
	olds, err := client.Str().GetSet("aa", "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		panic(err)
	}
	fmt.Println(olds)

	// range
	ranges, err := client.Str().GetRange("aa", 1, 35)
	if err != nil {
		panic(err)
	}
	fmt.Println(ranges)

	bs, err := client.Str().GetBit("aa", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(bs)



}
