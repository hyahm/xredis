package xredis

import "time"

type typeList struct {
	*GoRedis
}


func (gr *GoRedis) List() *typeList {
	return &typeList{gr	}
}


func (gr *typeList) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	if  err := gr.Ping(); err != nil {
		return  nil, err
	}

	return gr.client.BLPop(timeout, keys...).Result()
}


func (gr *typeList) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	if  err := gr.Ping(); err != nil {
		return  nil, err
	}

	return gr.client.BRPop(timeout, keys...).Result()
}

func (gr *typeList) BRPopLPush(source, destination string, timeout time.Duration) (string, error) {
	// 从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  "", err
	}

	return gr.client.BRPopLPush(source, destination,timeout).Result()
}



func (gr *typeList) LIndex(key string, index int64) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 通过索引获取列表中的元素
	if  err := gr.Ping(); err != nil {
		return  "", err
	}

	return gr.client.LIndex(key, index).Result()
}

func (gr *typeList) LInsert(key, op string, pivot, value interface{}) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 在列表的元素前或者后插入元素
	if  err := gr.Ping(); err != nil {
		return  0, err
	}

	return gr.client.LInsert(key, op, pivot, value).Result()
}

func (gr *typeList) LLen(key string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 获取列表长度
	if  err := gr.Ping(); err != nil {
		return  0, err
	}

	return gr.client.LLen(key).Result()
}

func (gr *typeList) LPop(key string) (string, error) {
	// 移出并获取列表的第一个元素
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  "", err
	}

	return gr.client.LPop(key).Result()
}

func (gr *typeList) LPush(key string, values ...interface{}) (int64, error) {
	// 将一个或多个值插入到列表头部
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  0, err
	}

	return gr.client.LPush(key, values...).Result()
}

func (gr *typeList) LPushX(key string, values ...interface{}) (int64, error) {
	// 将一个值插入到已存在的列表头部
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  0, err
	}

	return gr.client.LPushX(key, values...).Result()
}

func (gr *typeList) LRange(key string, start, stop int64) ([]string, error) {
	// 获取列表指定范围内的元素
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  nil, err
	}

	return gr.client.LRange(key, start, stop).Result()
}

func (gr *typeList) LRem(key string, count int64, value interface{}) (int64, error) {
	// 移除列表元素
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  0, err
	}

	return gr.client.LRem(key, count, value).Result()
}

func (gr *typeList) LSet(key string, index int64, value interface{}) (string, error) {
	// 通过索引设置列表元素的值
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.LSet(key, index, value).Result()
}

func (gr *typeList) LTrim(key string, start, stop int64) (string, error) {
	// 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.LTrim(key, start, stop).Result()
}

func (gr *typeList) RPop(key string) (string, error) {
	// 移除列表的最后一个元素，返回值为移除的元素。
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.RPop(key).Result()
}


func (gr *typeList) RPopLPush(source, destination string) (string, error) {
	// 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.RPopLPush(source, destination).Result()
}

func (gr *typeList) RPush(key string, values ...interface{}) (int64, error) {
	// 在列表中添加一个或多个值
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.RPush(key, values...).Result()
}

func (gr *typeList) RPushX(key string, values ...interface{}) (int64, error) {
	// 为已存在的列表添加值
	if gr == nil {
		panic("please conn first")
	}
	if err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.RPushX(key, values...).Result()
}