package xredis

import (
	"github.com/go-redis/redis"
)

type GoRedis struct {
	conf   *redis.Options
	client *redis.Client
}

func Conn(opt *redis.Options) (*GoRedis, error) {

	client := redis.NewClient(opt)
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	gr := &GoRedis{
		conf:   opt,
		client: client,
	}

	return gr, nil
}

func (gr *GoRedis) Ping() error {
	return gr.client.Ping().Err()
}

func (gr *GoRedis) Echo(message interface{}) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.Echo(message).Result()
}

func (gr *GoRedis) Quit() (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.Quit().Result()
}
