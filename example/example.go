package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/hyahm/xredis"
	"time"
)

type user struct {
	Name string
	Age int
	Id int64
	Role int
	Token string
}

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
	s := client.NewStr()
	if _, err := s.Set("aa", "bbb", time.Second*10); err != nil {
		panic(err)
	}
	fmt.Println("设置成功")

	value, err := s.Get("aa")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	//
	// 添加映射的缓存表
	tt  := client.NewTable()
	err = tt.InsertTemplate("user", user{})
	if err != nil {
		panic(err)
	}

	err = tt.SetPrefix("user", "www")
	if err != nil {
		panic(err)
	}
	// range
	fmt.Printf("%v", tt.GetKeys("user"))




}

