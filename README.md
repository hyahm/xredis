# xredis
redis

### 原生的驱动， 全部都放在了一起， 太乱， 将方法分类

如果技术允许的话，封装自定义一种高级缓存 --- 多key一值

例子
```go
package main

import (
	"fmt"
	"log"

	"github.com/hyahm/xredis"
)

func main() {

	conf := xredis.Options{
		Network:            "",
		Addr:               "172.28.147.151:6379",
		Dialer:             nil,
		Password:           "",
		DB:                 0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	}
	client, err := xredis.Conn(conf)
	if err != nil {
		panic(err)
	}

	if _, err := client.NewStr().Set("aa", "bbb"); err != nil {
		panic(err)
	}
	// lset  list+set 有序集合, 新增特殊有序集合
	ls := client.NewListSet()
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
	fmt.Println("设置成功")

}

```

