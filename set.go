package xredis

import "github.com/gomodule/redigo/redis"

type TypeSet struct {
	*GoRedis
	redis.Conn
}

func (gr *GoRedis) NewSet() *TypeSet {
	conn := gr.client.Get()
	return &TypeSet{gr, conn}
}

func (ts *TypeSet) SAdd(key string, members ...interface{}) (int64, error) {
	args := []interface{}{key}
	args = append(args, members...)
	return redis.Int64(ts.Do("sadd", args...))
}

func (ts *TypeSet) SCard(key string) (int64, error) {
	return redis.Int64(ts.Do("SCard", key))
}

func (ts *TypeSet) SDiff(key ...interface{}) ([]string, error) {
	return redis.Strings(ts.Do("SDiff", key...))
}
func (ts *TypeSet) SDiffStore(destination string, keys ...interface{}) (int64, error) {
	args := []interface{}{destination}
	args = append(args, keys...)
	return redis.Int64(ts.Do("SDiffStore", args...))
}

func (ts *TypeSet) SInter(keys ...interface{}) ([]string, error) {
	return redis.Strings(ts.Do("SInter", keys...))
}

func (ts *TypeSet) SInterStore(destination string, keys ...interface{}) (int64, error) {
	args := []interface{}{destination}
	args = append(args, keys...)
	return redis.Int64(ts.Do("SInterStore", args...))
}

func (ts *TypeSet) SIsMember(key string, member interface{}) (bool, error) {
	return redis.Bool(ts.Do("sismember", key, member))
}

func (ts *TypeSet) SMembers(key string) ([]string, error) {
	return redis.Strings(ts.Do("smembers", key))
}

func (ts *TypeSet) SMove(source, destination string, member interface{}) (bool, error) {
	return redis.Bool(ts.Do("SMove", source, destination, member))
}

func (ts *TypeSet) SPop(key string) (string, error) {
	return redis.String(ts.Do("SPop", key))
}

func (ts *TypeSet) SRandMember(key string) (string, error) {
	return redis.String(ts.Do("SRandMember", key))
}

func (ts *TypeSet) SRem(key string, members ...interface{}) (int64, error) {
	args := []interface{}{key}
	args = append(args, members...)
	return redis.Int64(ts.Do("SRem", args...))
}

func (ts *TypeSet) SUnion(keys ...interface{}) ([]string, error) {
	return redis.Strings(ts.Do("SUnion", keys...))
}

func (ts *TypeSet) SUnionStore(destination string, keys ...interface{}) (int64, error) {
	args := []interface{}{destination}
	args = append(args, keys...)
	return redis.Int64(ts.Do("SUnionStore", args...))
}

// func (ts *TypeSet) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
// 	return redis.Strings(ts.Do("SScan", key, cursor, match, count))
// }
