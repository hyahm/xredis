package main

import (
	"fmt"
	"log"
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
		Addr:               "192.168.50.250:6379",
		Dialer:             nil,
		Password:           "",
		DB:                 5,
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
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
	if !ok {
		_, err := set.SAdd("u5_download", "bbbb")
		if err != nil {
			log.Fatal(err)
		}
	}
	err = set.Expire("u5_download", 60*2)
	fmt.Println(err)

	// lset  list+set 有序集合, 新增特殊有序集合
	ls := rconn.NewListSet()
	//
	ls.LPush("aaa", "bbb")
	ls.LPush("aaa", "ccc")
	ls.LPush("aaa", "ddd")
	ls.LPush("aaa", "eee")
	ls.LPush("aaa", "eee")
	result, err := ls.Range("aaa", 0, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result) //  [eee ddd ccc bbb]
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
