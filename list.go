package xredis

import (
	"github.com/gomodule/redigo/redis"
)

type TypeList struct {
	*GoRedis
	redis.Conn
}

func (gr *GoRedis) NewList() *TypeList {
	conn := gr.client.Get()
	return &TypeList{gr, conn}
}

// func (tl *TypeList) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
// 	return redis.Strings(tl.Do("BLPop",))
// }

// func (tl *TypeList) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
// 	// 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
// 	return tl.client.BRPop(timeout, keys...).Result()
// }

// func (tl *TypeList) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
// 	// 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
// 	return tl.client.BRPopLPush(source, destination, timeout).Result()
// }

func (tl *TypeList) LIndex(key string, index int64) (string, error) {
	// 通过索引获取列表中的元素
	return redis.String(tl.Do("lindex", key, index))
}

func (tl *TypeList) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	// 在列表的元素前或者后插入元素
	return redis.Int64(tl.Do("LInsert", key, op, pivot, value))
}

func (tl *TypeList) LLen(key string) (int64, error) {
	// 获取列表长度
	return redis.Int64(tl.Do("LLen", key))
}

func (tl *TypeList) LPop(key string) (string, error) {
	// 移出并获取列表的第一个元素
	return redis.String(tl.Do("lpop", key))
}

func (tl *TypeList) LPush(key string, values ...interface{}) (int64, error) {
	// 将一个或多个值插入到列表头部
	args := []interface{}{key}
	args = append(args, values...)
	return redis.Int64(tl.Do("lpush", args...))
}

func (tl *TypeList) LPushX(key string, value interface{}) (int64, error) {
	// 将一个值插入到已存在的列表头部
	return redis.Int64(tl.Do("LPushX", key, value))
}

func (tl *TypeList) LRange(key string, start, stop int64) ([]string, error) {
	// 获取列表指定范围内的元素
	return redis.Strings(tl.Do("LRange", key, start, stop))
}

func (tl *TypeList) LRem(key string, count int64, value interface{}) (int64, error) {
	// 移除列表元素
	return redis.Int64(tl.Do("LRem", key, count, value))
}

func (tl *TypeList) LSet(key string, index int64, value interface{}) (string, error) {
	// 通过索引设置列表元素的值
	return redis.String(tl.Do("LSet", key, index, value))
}

func (tl *TypeList) LTrim(key string, start, stop int64) (string, error) {
	// 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
	return redis.String(tl.Do("LTrim", key, start, stop))
}

func (tl *TypeList) RPop(key string) (string, error) {
	// 移除列表的最后一个元素，返回值为移除的元素。
	return redis.String(tl.Do("RPop", key))
}

func (tl *TypeList) RPopLPush(source, destination string) (string, error) {
	// 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	return redis.String(tl.Do("RPopLPush", source, destination))
}

func (tl *TypeList) RPush(key string, values ...interface{}) (int64, error) {
	// 在列表中添加一个或多个值
	args := []interface{}{key}
	args = append(args, values...)
	return redis.Int64(tl.Do("RPush", args...))

}

func (tl *TypeList) RPushX(key string, value interface{}) (int64, error) {
	// 为已存在的列表添加值
	return redis.Int64(tl.Do("RPushX", key, value))
}
