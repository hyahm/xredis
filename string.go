package xredis

import (
	"time"
)

type typeString struct {
	*GoRedis
}


func (gr *GoRedis) Str() *typeString {
	return &typeString{gr}
}


func (gr *typeString) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 设置指定 key 的值
	if  err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.Set(key, value, expiration).Result()
}

func (gr *typeString) Get(key string) (string,error) {
	if gr == nil {
		panic("please conn first")
	}
	// 获取指定 key 的值。
	if  err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.Get(key).Result()
}

func (gr *typeString) GetRange(key string, start, end int64) (string,error) {
	if gr == nil {
		panic("please conn first")
	}
	// 返回 key 中字符串值的子字符
	if  err := gr.Ping(); err != nil {
		return "",err
	}
	return gr.client.GetRange(key, start, end).Result()
}

func (gr *typeString) GetSet(key string, value interface{}) (string,error) {
	// 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return "",err
	}
	return gr.client.GetSet(key, value).Result()
}


func (gr *typeString) GetBit(key string, offset int64) (int64,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0,err
	}
	return gr.client.GetBit(key, offset).Result()
}


func (gr *typeString) SetBit(key string, offset int64, value int) (int64,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0,err
	}
	return gr.client.SetBit(key, offset, value).Result()
}
func (gr *typeString) SetEx(key string, expiration time.Duration) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return false, err
	}
	return gr.client.Expire(key,expiration).Result()
}
func (gr *typeString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return false, err
	}
	return gr.client.SetNX(key , value ,expiration).Result()
}
func (gr *typeString) MGet(keys ...string) ([]interface{},error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil,err
	}
	return gr.client.MGet(keys...).Result()
}
func (gr *typeString) SetRange(key string, offset int64, value string) error {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return err
	}
	return gr.client.SetRange(key,offset,value).Err()
}
func (gr *typeString) StrLen(key string) (int64,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0,err
	}
	return gr.client.StrLen(key).Result()
}
func (gr *typeString) MSet(pairs ...interface{}) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return "" ,err
	}
	return gr.client.MSet(pairs...).Result()
}
func (gr *typeString) MSetNX(pairs ...interface{}) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return false, err
	}
	return gr.client.MSetNX(pairs...).Result()
}
func (gr *typeString) Incr(key string) (int64, error)  {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0 ,err
	}
	return gr.client.Incr(key).Result()
}
func (gr *typeString) IncrBy(key string, value int64) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}
	return gr.client.IncrBy(key, value).Result()
}
func (gr *typeString) IncrByFloat(key string, value float64) (float64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}
	return gr.client.IncrByFloat(key, value).Result()
}

func (gr *typeString) Decr(key string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}
	return gr.client.Decr(key).Result()
}

func (gr *typeString) DecrBy(key string, value int64) (int64, error)  {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}
	return gr.client.DecrBy(key, value).Result()
}

func (gr *typeString) Append(key string, value string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}
	return gr.client.Append(key, value).Result()
}
