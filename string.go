package xredis

import (
	"time"
)

type typeString struct {
	*GoRedis
}


func (gr *GoRedis) NewStr() *typeString {
	return &typeString{gr}
}


func (tsr *typeString) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	if tsr == nil {
		panic("please conn first")
	}
	// 设置指定 key 的值
	if  err := tsr.Ping(); err != nil {
		return "", err
	}

	return tsr.client.Set(key, value, expiration).Result()
}

func (tsr *typeString) Get(key string) (string,error) {
	if tsr == nil {
		panic("please conn first")
	}
	// 获取指定 key 的值。
	if  err := tsr.Ping(); err != nil {
		return "", err
	}

	return tsr.client.Get(key).Result()
}

func (tsr *typeString) GetRange(key string, start, end int64) (string,error) {
	if tsr == nil {
		panic("please conn first")
	}
	// 返回 key 中字符串值的子字符
	if  err := tsr.Ping(); err != nil {
		return "",err
	}
	return tsr.client.GetRange(key, start, end).Result()
}

func (tsr *typeString) GetSet(key string, value interface{}) (string,error) {
	// 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return "",err
	}
	return tsr.client.GetSet(key, value).Result()
}


func (tsr *typeString) GetBit(key string, offset int64) (int64,error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.client.GetBit(key, offset).Result()
}


func (tsr *typeString) SetBit(key string, offset int64, value int) (int64,error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.client.SetBit(key, offset, value).Result()
}
func (tsr *typeString) SetEx(key string, expiration time.Duration) (bool, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.client.Expire(key,expiration).Result()
}
func (tsr *typeString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.client.SetNX(key , value ,expiration).Result()
}
func (tsr *typeString) MGet(keys ...string) ([]interface{},error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return nil,err
	}
	return tsr.client.MGet(keys...).Result()
}
func (tsr *typeString) SetRange(key string, offset int64, value string) error {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return err
	}
	return tsr.client.SetRange(key,offset,value).Err()
}
func (tsr *typeString) StrLen(key string) (int64,error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.client.StrLen(key).Result()
}
func (tsr *typeString) MSet(pairs ...interface{}) (string, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return "" ,err
	}
	return tsr.client.MSet(pairs...).Result()
}
func (tsr *typeString) MSetNX(pairs ...interface{}) (bool, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.client.MSetNX(pairs...).Result()
}
func (tsr *typeString) Incr(key string) (int64, error)  {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0 ,err
	}
	return tsr.client.Incr(key).Result()
}
func (tsr *typeString) IncrBy(key string, value int64) (int64, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.IncrBy(key, value).Result()
}
func (tsr *typeString) IncrByFloat(key string, value float64) (float64, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.IncrByFloat(key, value).Result()
}

func (tsr *typeString) Decr(key string) (int64, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.Decr(key).Result()
}

func (tsr *typeString) DecrBy(key string, value int64) (int64, error)  {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.DecrBy(key, value).Result()
}

func (tsr *typeString) Append(key string, value string) (int64, error) {
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.Append(key, value).Result()
}
