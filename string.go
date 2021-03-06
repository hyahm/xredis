package xredis

import (
	"github.com/gomodule/redigo/redis"
)

type TypeString struct {
	*GoRedis
	redis.Conn
}

func (gr *GoRedis) NewStr() *TypeString {
	conn := gr.client.Get()
	return &TypeString{gr, conn}
}

func (tsr *TypeString) Set(key string, value interface{}) (string, error) {
	return redis.String(tsr.Do("set", key, value))
}

func (tsr *TypeString) Get(key string) (string, error) {
	return redis.String(tsr.Do("Get", key))
}

func (tsr *TypeString) GetRange(key string, start, end int64) (string, error) {
	return redis.String(tsr.Do("GetRange", key, start, end))
}

func (tsr *TypeString) GetSet(key string, value interface{}) (string, error) {
	// 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
	return redis.String(tsr.Do("GetSet", key, value))
}

// func (tsr *TypeString) GetBit(key string, offset int64) (int64, error) {
// 	return tsr.client.GetBit(key, offset).Result()
// }

// func (tsr *TypeString) SetBit(key string, offset int64, value int) (int64, error) {
// 	return tsr.client.SetBit(key, offset, value).Result()
// }
// func (tsr *TypeString) SetEx(key string, expiration time.Duration) (bool, error) {
// 	return tsr.client.Expire(key, expiration).Result()
// }
// func (tsr *TypeString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
// 	return tsr.client.SetNX(key, value, expiration).Result()
// }
// func (tsr *TypeString) MGet(keys ...string) ([]interface{}, error) {
// 	return tsr.client.MGet(keys...).Result()
// }
// func (tsr *TypeString) SetRange(key string, offset int64, value string) error {
// 	return tsr.client.SetRange(key, offset, value).Err()
// }
// func (tsr *TypeString) StrLen(key string) (int64, error) {
// 	return tsr.client.StrLen(key).Result()
// }
// func (tsr *TypeString) MSet(pairs ...interface{}) (string, error) {
// 	return tsr.client.MSet(pairs...).Result()
// }
// func (tsr *TypeString) MSetNX(pairs ...interface{}) (bool, error) {
// 	return tsr.client.MSetNX(pairs...).Result()
// }
// func (tsr *TypeString) Incr(key string) (int64, error) {
// 	return tsr.client.Incr(key).Result()
// }
// func (tsr *TypeString) IncrBy(key string, value int64) (int64, error) {
// 	return tsr.client.IncrBy(key, value).Result()
// }
// func (tsr *TypeString) IncrByFloat(key string, value float64) (float64, error) {
// 	return tsr.client.IncrByFloat(key, value).Result()
// }

// func (tsr *TypeString) Decr(key string) (int64, error) {
// 	return tsr.client.Decr(key).Result()
// }

// func (tsr *TypeString) DecrBy(key string, value int64) (int64, error) {
// 	return tsr.client.DecrBy(key, value).Result()
// }

// func (tsr *TypeString) Append(key string, value string) (int64, error) {
// 	return tsr.client.Append(key, value).Result()
// }
