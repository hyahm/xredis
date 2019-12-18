package xredis

import "time"

type typeList struct {
	tl *GoRedis
}


func (gr *GoRedis) NewList() *typeList {
	return &typeList{gr	}
}


func (tl *typeList) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	if tl.tl == nil {
		panic("please conn first")
	}
	// 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	if  err := tl.tl.Ping(); err != nil {
		return  nil, err
	}

	return tl.tl.client.BLPop(timeout, keys...).Result()
}


func (tl *typeList) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	if tl.tl == nil {
		panic("please conn first")
	}
	// 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	if  err := tl.tl.Ping(); err != nil {
		return  nil, err
	}

	return tl.tl.client.BRPop(timeout, keys...).Result()
}

func (tl *typeList) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	// 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	if tl.tl == nil {
		panic("please conn first")
	}
	if  err := tl.tl.Ping(); err != nil {
		return  "", err
	}

	return tl.tl.client.BRPopLPush(source, destination,timeout).Result()
}



func (tl *typeList) LIndex(key string, index int64) (string, error) {
	if tl.tl == nil {
		panic("please conn first")
	}
	// 通过索引获取列表中的元素
	if  err := tl.tl.Ping(); err != nil {
		return  "", err
	}

	return tl.tl.client.LIndex(key, index).Result()
}

func (tl *typeList) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	if tl.tl == nil {
		panic("please conn first")
	}
	// 在列表的元素前或者后插入元素
	if  err := tl.tl.Ping(); err != nil {
		return  0, err
	}

	return tl.tl.client.LInsert(key, op, pivot, value).Result()
}

func (tl *typeList) LLen(key string) (int64, error) {
	if tl.tl == nil {
		panic("please conn first")
	}
	// 获取列表长度
	if  err := tl.tl.Ping(); err != nil {
		return  0, err
	}

	return tl.tl.client.LLen(key).Result()
}

func (tl *typeList) LPop(key string) (string, error) {
	// 移出并获取列表的第一个元素
	if tl.tl == nil {
		panic("please conn first")
	}
	if  err := tl.tl.Ping(); err != nil {
		return  "", err
	}

	return tl.tl.client.LPop(key).Result()
}

func (tl *typeList) LPush(key string, values ...interface{}) (int64, error) {
	// 将一个或多个值插入到列表头部
	if tl.tl == nil {
		panic("please conn first")
	}
	if  err := tl.tl.Ping(); err != nil {
		return  0, err
	}

	return tl.tl.client.LPush(key, values...).Result()
}

func (tl *typeList) LPushX(key string, values ...interface{}) (int64, error) {
	// 将一个值插入到已存在的列表头部
	if tl.tl == nil {
		panic("please conn first")
	}
	if  err := tl.tl.Ping(); err != nil {
		return  0, err
	}

	return tl.tl.client.LPushX(key, values...).Result()
}

func (tl *typeList) LRange(key string, start, stop int64) ([]string, error) {
	// 获取列表指定范围内的元素
	if tl.tl == nil {
		panic("please conn first")
	}
	if  err := tl.tl.Ping(); err != nil {
		return  nil, err
	}

	return tl.tl.client.LRange(key, start, stop).Result()
}

func (tl *typeList) LRem(key string, count int64, value interface{}) (int64, error) {
	// 移除列表元素
	if tl.tl == nil {
		panic("please conn first")
	}
	if  err := tl.tl.Ping(); err != nil {
		return  0, err
	}

	return tl.tl.client.LRem(key, count, value).Result()
}

func (tl *typeList) LSet(key string, index int64, value interface{}) (string, error) {
	// 通过索引设置列表元素的值
	if tl.tl == nil {
		panic("please conn first")
	}
	if err := tl.tl.Ping(); err != nil {
		return "", err
	}

	return tl.tl.client.LSet(key, index, value).Result()
}

func (tl *typeList) LTrim(key string, start, stop int64) (string, error) {
	// 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
	if tl.tl == nil {
		panic("please conn first")
	}
	if err := tl.tl.Ping(); err != nil {
		return "", err
	}

	return tl.tl.client.LTrim(key, start, stop).Result()
}

func (tl *typeList) RPop(key string) (string, error) {
	// 移除列表的最后一个元素，返回值为移除的元素。
	if tl.tl == nil {
		panic("please conn first")
	}
	if err := tl.tl.Ping(); err != nil {
		return "", err
	}

	return tl.tl.client.RPop(key).Result()
}


func (tl *typeList) RPopLPush(source, destination string) (string, error) {
	// 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	if tl.tl == nil {
		panic("please conn first")
	}
	if err := tl.tl.Ping(); err != nil {
		return "", err
	}

	return tl.tl.client.RPopLPush(source, destination).Result()
}

func (tl *typeList) RPush(key string, values ...interface{}) (int64, error) {
	// 在列表中添加一个或多个值
	if tl.tl == nil {
		panic("please conn first")
	}
	if err := tl.tl.Ping(); err != nil {
		return 0, err
	}

	return tl.tl.client.RPush(key, values...).Result()
}

func (tl *typeList) RPushX(key string, values ...interface{}) (int64, error) {
	// 为已存在的列表添加值
	if tl.tl == nil {
		panic("please conn first")
	}
	if err := tl.tl.Ping(); err != nil {
		return 0, err
	}

	return tl.tl.client.RPushX(key, values...).Result()
}