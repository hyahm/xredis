package xredis

import (
	"time"
)

type TypeString struct {
	tsr *GoRedis
}


func (gr *GoRedis) NewStr() *TypeString {
	return &TypeString{gr}
}


func (tsr *TypeString) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	// 设置指定 key 的值
	if  err := tsr.tsr.Ping(); err != nil {
		return "", err
	}

	return tsr.tsr.client.Set(key, value, expiration).Result()
}

func (tsr *TypeString) Get(key string) (string,error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	// 获取指定 key 的值。
	if  err := tsr.tsr.Ping(); err != nil {
		return "", err
	}

	return tsr.tsr.client.Get(key).Result()
}

func (tsr *TypeString) GetRange(key string, start, end int64) (string,error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	// 返回 key 中字符串值的子字符
	if  err := tsr.tsr.Ping(); err != nil {
		return "",err
	}
	return tsr.tsr.client.GetRange(key, start, end).Result()
}

func (tsr *TypeString) GetSet(key string, value interface{}) (string,error) {
	// 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return "",err
	}
	return tsr.tsr.client.GetSet(key, value).Result()
}


func (tsr *TypeString) GetBit(key string, offset int64) (int64,error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.tsr.client.GetBit(key, offset).Result()
}


func (tsr *TypeString) SetBit(key string, offset int64, value int) (int64,error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.tsr.client.SetBit(key, offset, value).Result()
}
func (tsr *TypeString) SetEx(key string, expiration time.Duration) (bool, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.tsr.client.Expire(key,expiration).Result()
}
func (tsr *TypeString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.tsr.client.SetNX(key , value ,expiration).Result()
}
func (tsr *TypeString) MGet(keys ...string) ([]interface{},error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return nil,err
	}
	return tsr.tsr.client.MGet(keys...).Result()
}
func (tsr *TypeString) SetRange(key string, offset int64, value string) error {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return err
	}
	return tsr.tsr.client.SetRange(key,offset,value).Err()
}
func (tsr *TypeString) StrLen(key string) (int64,error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.tsr.client.StrLen(key).Result()
}
func (tsr *TypeString) MSet(pairs ...interface{}) (string, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return "" ,err
	}
	return tsr.tsr.client.MSet(pairs...).Result()
}
func (tsr *TypeString) MSetNX(pairs ...interface{}) (bool, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.tsr.client.MSetNX(pairs...).Result()
}
func (tsr *TypeString) Incr(key string) (int64, error)  {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0 ,err
	}
	return tsr.tsr.client.Incr(key).Result()
}
func (tsr *TypeString) IncrBy(key string, value int64) (int64, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.tsr.client.IncrBy(key, value).Result()
}
func (tsr *TypeString) IncrByFloat(key string, value float64) (float64, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.tsr.client.IncrByFloat(key, value).Result()
}

func (tsr *TypeString) Decr(key string) (int64, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.tsr.client.Decr(key).Result()
}

func (tsr *TypeString) DecrBy(key string, value int64) (int64, error)  {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.tsr.client.DecrBy(key, value).Result()
}

func (tsr *TypeString) Append(key string, value string) (int64, error) {
	if tsr.tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.tsr.client.Append(key, value).Result()
}
