package xredis

import (
	"github.com/go-redis/redis/v7"
)



type GoRedis struct {
	conf *redis.Options
	client *redis.Client
	//Err error
}

func Conn(opt *redis.Options) (*GoRedis, error) {

	client := redis.NewClient(opt)
	if err := client.Ping().Err();err != nil {
		return nil, err
	}
	gr := &GoRedis{
		conf:   opt,
		client: client,
	}

	return gr, nil
}

func (gr *GoRedis) ping() error {
	if err := gr.client.Ping().Err(); err != nil {
		// 重连一次
		gr.client = redis.NewClient(gr.conf)
		if err := gr.client.Ping().Err();err != nil {
			return err
		}
	}
	return  nil
}

