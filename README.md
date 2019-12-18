# xredis
redis

### 原生的驱动， 全部都放在了一起， 太乱， 将方法分类
感觉是体力活， 等待堆体力

如果技术允许的话，封装自定义一种高级缓存 --- 多key一值

例子
```go
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

}
```

