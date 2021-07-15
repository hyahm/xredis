package xredis

import "github.com/gomodule/redigo/redis"

type TypeSet struct {
	redis.Conn
}

func (gr *GoRedis) NewSet() *TypeSet {
	conn := gr.client.Get()
	return &TypeSet{conn}
}

func (ts *TypeSet) SAdd(key string, members ...interface{}) (int64, error) {
	args := make([]interface{}, 1, len(members)+1)
	args[0] = key
	args = append(args, members...)
	return redis.Int64(ts.Do("sadd", args...))
}

// func (ts *TypeSet) SCard(key string) (int64, error) {
// 	return ts.client.SCard(key).Result()
// }

// func (ts *TypeSet) SDiff(key ...string) ([]string, error) {
// 	return ts.client.SDiff(key...).Result()
// }
// func (ts *TypeSet) SDiffStore(destination string, keys ...string) (int64, error) {
// 	return ts.client.SDiffStore(destination, keys...).Result()
// }

// func (ts *TypeSet) SInter(keys ...string) ([]string, error) {
// 	return ts.client.SInter(keys...).Result()
// }

// func (ts *TypeSet) SInterStore(destination string, keys ...string) (int64, error) {
// 	return ts.client.SInterStore(destination, keys...).Result()
// }

func (ts *TypeSet) SIsMember(key string, member interface{}) (bool, error) {
	return redis.Bool(ts.Do("sismember", key, member))
}

func (ts *TypeSet) SMembers(key string) ([]string, error) {
	return redis.Strings(ts.Do("smembers", key))
}

// func (ts *TypeSet) SMove(source, destination string, member interface{}) (bool, error) {
// 	return ts.client.SMove(source, destination, member).Result()
// }

// func (ts *TypeSet) SPop(key string) (string, error) {
// 	return ts.client.SPop(key).Result()
// }

// func (ts *TypeSet) SRandMember(key string) (string, error) {
// 	return ts.client.SRandMember(key).Result()
// }

// func (ts *TypeSet) SRandMemberN(key string, count int64) ([]string, error) {
// 	return ts.client.SRandMemberN(key, count).Result()
// }

// func (ts *TypeSet) SRem(key string, members ...interface{}) (int64, error) {
// 	return ts.client.SRem(key, members...).Result()
// }

// func (ts *TypeSet) SUnion(keys ...string) ([]string, error) {
// 	return ts.client.SUnion(keys...).Result()
// }

// func (ts *TypeSet) SUnionStore(destination string, keys ...string) (int64, error) {
// 	return ts.client.SUnionStore(destination, keys...).Result()
// }

// func (ts *TypeSet) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
// 	return ts.client.SScan(key, cursor, match, count).Result()
// }
