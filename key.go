package xredis

import (
	"github.com/gomodule/redigo/redis"
)

func (gr *GoRedis) Del(keys ...interface{}) (int64, error) {
	conn := gr.client.Get()
	return redis.Int64(conn.Do("del", keys...))
}

func (gr *GoRedis) Dump(key string) (string, error) {
	conn := gr.client.Get()
	return redis.String(conn.Do("dump", key))
}

func (gr *GoRedis) RPushX(key string, value interface{}) (int64, error) {
	// 为已存在的列表添加值
	conn := gr.client.Get()
	return redis.Int64(conn.Do("RPushX", key, value))
}

func (gr *GoRedis) ExpireAt(key string, time int64) (bool, error) {
	conn := gr.client.Get()
	return redis.Bool(conn.Do("expireAt", key, time))
}

func (gr *GoRedis) Keys(pattern string) ([]string, error) {
	conn := gr.client.Get()
	return redis.Strings(conn.Do("keys", pattern))
}

func (gr *GoRedis) Move(key string, db int64) (bool, error) {
	conn := gr.client.Get()
	return redis.Bool(conn.Do("move", db))
}

func (gr *GoRedis) Persist(key string) (bool, error) {
	conn := gr.client.Get()
	return redis.Bool(conn.Do("persist", key))
}

func (gr *GoRedis) Expire(key string, second int) error {
	conn := gr.client.Get()
	_, err := redis.Int(conn.Do("EXPIRE", key, second))
	return err
}

func (gr *GoRedis) GetTTL(key string) (int, error) {
	conn := gr.client.Get()
	return redis.Int(conn.Do("ttl", key))
}

func (gr *GoRedis) RandomKey() (string, error) {
	conn := gr.client.Get()
	return redis.String(conn.Do("rename"))
}

func (gr *GoRedis) Rename(key, newkey string) (string, error) {
	conn := gr.client.Get()
	return redis.String(conn.Do("rename", key, newkey))
}

func (gr *GoRedis) RenameNX(key, newkey string) (bool, error) {
	conn := gr.client.Get()
	return redis.Bool(conn.Do("RenameNX", key, newkey))
}

func (gr *GoRedis) Type(key string) (string, error) {
	conn := gr.client.Get()
	return redis.String(conn.Do("type", key))
}
