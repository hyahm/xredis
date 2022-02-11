package xredis

import "github.com/gomodule/redigo/redis"

type Position string

const BEFORE Position = "BEFORE"
const AFTER Position = "AFTER"

const DEFAULTlSET string = "list_for_"

// 新封装的数据格式， list + set 组合而成的有序 set
// 注意： 数据类型会 list类型（名字前面会有个list_for_）和 set 类型
type TypeListSet struct {
	*GoRedis
	redis.Conn
}

func (gr *GoRedis) NewListSet() *TypeListSet {
	conn := gr.client.Get()
	return &TypeListSet{gr, conn}
}

// 查看元素是否存在
func (tls *TypeListSet) IsMembers(key string, value interface{}) (bool, error) {

	return tls.NewSet().SIsMember(key, value)
}

func (tls *TypeListSet) Index(key string, index int64) (string, error) {
	// 通过索引获取列表中的元素
	return redis.String(tls.Do("lindex", DEFAULTlSET+key, index))
}

// 返回 int64: 如果 为 -1 就是存在这个key
func (tls *TypeListSet) Insert(key string, op Position, pivot, value interface{}) (int64, error) {
	// 在列表的某元素前或者后插入元素
	// 验证是否存在key
	ok, err := tls.NewSet().SIsMember(key, value)
	if err != nil {
		return 0, err
	}
	if ok {
		return -1, nil
	}
	n, err := tls.NewSet().SAdd(key, value)
	if err != nil {
		return n, err
	}
	if n == 0 {
		return -1, nil
	}
	return redis.Int64(tls.Do("LInsert", DEFAULTlSET+key, op, pivot, value))
}

func (tls *TypeListSet) Len(key string) (int64, error) {
	// 获取列表长度
	return redis.Int64(tls.Do("LLen", DEFAULTlSET+key))
}

func (tls *TypeListSet) LPop(key string) (string, error) {
	// 移出并获取列表的第一个元素
	value, err := redis.String(tls.Do("lpop", DEFAULTlSET+key))
	if err != nil {
		return value, err
	}
	tls.NewSet().SRem(key, value)
	return value, err
}

// 将一个或多个值插入到列表头部
func (tls *TypeListSet) LPush(key string, values ...interface{}) (int64, error) {

	var num int64 = 0
	for _, v := range values {
		ok, err := tls.NewSet().SIsMember(key, v)
		if err != nil {
			return 0, err
		}
		if ok {
			continue
		}
		n, err := tls.NewSet().SAdd(key, v)
		if err != nil {
			return n, err
		}
		if n == 0 {
			return -1, nil
		}
		_, err = redis.Int64(tls.Do("lpush", DEFAULTlSET+key, v))
		if err != nil {
			return 0, err
		}
		num++
	}

	return num, nil
}

func (tls *TypeListSet) RPop(key string) (string, error) {
	// 移除列表的最后一个元素，返回值为移除的元素。
	value, err := redis.String(tls.Do("RPop", DEFAULTlSET+key))
	if err != nil {
		return value, err
	}
	tls.NewSet().SRem(key, value)
	return value, err
}

func (tls *TypeListSet) Range(key string, start, stop int64) ([]string, error) {
	// 获取列表指定范围内的元素
	return redis.Strings(tls.Do("LRange", DEFAULTlSET+key, start, stop))
}

// 在列表后添加一个或多个值
func (tls *TypeListSet) RPush(key string, values ...interface{}) (int64, error) {

	var num int64 = 0
	for _, v := range values {
		ok, err := tls.NewSet().SIsMember(key, v)
		if err != nil {
			return 0, err
		}
		if ok {
			continue
		}
		n, err := tls.NewSet().SAdd(key, v)
		if err != nil {
			return n, err
		}
		if n == 0 {
			return -1, nil
		}
		_, err = redis.Int64(tls.Do("RPush", DEFAULTlSET+key, v))
		if err != nil {
			return 0, err
		}
		num++
	}
	return num, nil
}
