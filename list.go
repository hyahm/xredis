package xredis

import "time"

type TypeList struct {
	*GoRedis
}

func (gr *GoRedis) NewList() *TypeList {
	return &TypeList{gr}
}

func (tl *TypeList) Del(keys ...string) (int64, error) {
	return tl.client.Del(keys...).Result()
}

func (tl *TypeList) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	return tl.client.BLPop(timeout, keys...).Result()
}

func (tl *TypeList) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	// 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	return tl.client.BRPop(timeout, keys...).Result()
}

func (tl *TypeList) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	// 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	return tl.client.BRPopLPush(source, destination, timeout).Result()
}

func (tl *TypeList) LIndex(key string, index int64) (string, error) {
	// 通过索引获取列表中的元素
	return tl.client.LIndex(key, index).Result()
}

func (tl *TypeList) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	// 在列表的元素前或者后插入元素
	return tl.client.LInsert(key, op, pivot, value).Result()
}

func (tl *TypeList) LLen(key string) (int64, error) {
	// 获取列表长度
	return tl.client.LLen(key).Result()
}

func (tl *TypeList) LPop(key string) (string, error) {
	// 移出并获取列表的第一个元素
	return tl.client.LPop(key).Result()
}

func (tl *TypeList) LPush(key string, values ...interface{}) (int64, error) {
	// 将一个或多个值插入到列表头部
	return tl.client.LPush(key, values...).Result()
}

func (tl *TypeList) LPushX(key string, value interface{}) (int64, error) {
	// 将一个值插入到已存在的列表头部
	return tl.client.LPushX(key, value).Result()
}

func (tl *TypeList) LRange(key string, start, stop int64) ([]string, error) {
	// 获取列表指定范围内的元素
	return tl.client.LRange(key, start, stop).Result()
}

func (tl *TypeList) LRem(key string, count int64, value interface{}) (int64, error) {
	// 移除列表元素
	return tl.client.LRem(key, count, value).Result()
}

func (tl *TypeList) LSet(key string, index int64, value interface{}) (string, error) {
	// 通过索引设置列表元素的值
	return tl.client.LSet(key, index, value).Result()
}

func (tl *TypeList) LTrim(key string, start, stop int64) (string, error) {
	// 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
	return tl.client.LTrim(key, start, stop).Result()
}

func (tl *TypeList) RPop(key string) (string, error) {
	// 移除列表的最后一个元素，返回值为移除的元素。
	return tl.client.RPop(key).Result()
}

func (tl *TypeList) RPopLPush(source, destination string) (string, error) {
	// 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	return tl.client.RPopLPush(source, destination).Result()
}

func (tl *TypeList) RPush(key string, values ...interface{}) (int64, error) {
	// 在列表中添加一个或多个值
	return tl.client.RPush(key, values...).Result()
}

func (tl *TypeList) RPushX(key string, value interface{}) (int64, error) {
	// 为已存在的列表添加值
	return tl.client.RPushX(key, value).Result()
}
