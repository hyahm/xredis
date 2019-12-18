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

func (gr *GoRedis) Ping() error {
	if err := gr.client.Ping().Err(); err != nil {
		// 重连一次
		gr.client = redis.NewClient(gr.conf)
		if err := gr.client.Ping().Err();err != nil {
			return err
		}
	}
	return  nil
}

func (gr *GoRedis) Echo(message interface{}) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return "",  err
	}

	return gr.client.Echo(message).Result()
}

func (gr *GoRedis) Quit( ) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return "",  err
	}

	return gr.client.Quit().Result()
}
