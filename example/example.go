package main

import (
	"fmt"
	"time"

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
	conf := xredis.Options{
		Network:            "tcp",
		Addr:               "23.225.165.34:6379",
		Dialer:             nil,
		Password:           "hugonodahaiten",
		DB:                 5,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        5 * time.Second,
		WriteTimeout:       0,
		PoolSize:           10,
		MinIdleConns:       0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		ConnectTimeout:     3 * time.Second,
	}
	rconn, err := xredis.Conn(conf)
	if err != nil {
		fmt.Println(err)
	}
	set := rconn.NewSet()
	// ok, _ := set.SIsMember("u5_downloading", member)
	ok, err := set.Exists("u5_download")
	fmt.Println(ok)
	fmt.Println(err)

	// if _, err := client.NewStr().Set("aa", "bbb", time.Second*10); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("设置成功")

	// value, err := client.NewStr().Get("aa")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(value)

}
