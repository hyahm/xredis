package xredis

import "github.com/gomodule/redigo/redis"

type TypeHash struct {
	*GoRedis
	redis.Conn
}

func (gr *GoRedis) NewHash() *TypeHash {
	conn := gr.client.Get()
	return &TypeHash{gr, conn}
}

func (th *TypeHash) HDel(key string, fields ...string) (int64, error) {
	args := []interface{}{key}
	args = append(args, fields)
	return redis.Int64(th.Do("hdel", args))
}
func (th *TypeHash) HExists(key, field string) (bool, error) {
	return redis.Bool(th.Do("HExists", key, field))
}

func (th *TypeHash) HGet(key, field string) (string, error) {
	return redis.String(th.Do("HGet", key, field))
}

func (th *TypeHash) HGetAll(key string) (map[string]string, error) {
	return redis.StringMap(th.Do("HGetAll", key))
}
func (th *TypeHash) HIncrBy(key, field string, incr int64) (int64, error) {
	return redis.Int64(th.Do("HIncrBy", key, field, incr))
}
func (th *TypeHash) HIncrByFloat(key, field string, incr float64) (float64, error) {
	return redis.Float64(th.Do("HIncrByFloat", key, field, incr))
}
func (th *TypeHash) HKeys(key string) ([]string, error) {
	return redis.Strings(th.Do("HKeys", key))
}
func (th *TypeHash) HLen(key string) (int64, error) {
	return redis.Int64(th.Do("HLen", key))
}
func (th *TypeHash) HMGet(key string, fields ...interface{}) ([][]byte, error) {
	args := []interface{}{key}
	args = append(args, fields...)
	return redis.ByteSlices(th.Do("HMGet", args...))
}

func (th *TypeHash) HMSet(key string, fields map[string]interface{}) (string, error) {
	return redis.String(th.Do("HMSet", key, fields))
}
func (th *TypeHash) HSet(key, field string, value interface{}) (bool, error) {
	return redis.Bool(th.Do("HSet", key, field, value))
}

func (th *TypeHash) HSetNX(key, field string, value interface{}) (bool, error) {
	return redis.Bool(th.Do("HSetNX", key, field, value))
}

func (th *TypeHash) HVals(key string) ([]string, error) {
	return redis.Strings(th.Do("HVals", key))
}

// func (th *TypeHash) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
// 	return th.client.HScan(key, cursor, match, count).Result()
// }
