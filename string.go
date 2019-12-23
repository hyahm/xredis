package xredis

import (
	"time"
)

<<<<<<< HEAD
type typeString struct {
	*GoRedis
=======
type TypeString struct {
	tsr *GoRedis
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
}


func (gr *GoRedis) NewStr() *TypeString {
	return &TypeString{gr}
}


<<<<<<< HEAD
func (tsr *typeString) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	if tsr == nil {
=======
func (tsr *TypeString) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 设置指定 key 的值
	if  err := tsr.Ping(); err != nil {
		return "", err
	}

	return tsr.client.Set(key, value, expiration).Result()
}

<<<<<<< HEAD
func (tsr *typeString) Get(key string) (string,error) {
	if tsr == nil {
=======
func (tsr *TypeString) Get(key string) (string,error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 获取指定 key 的值。
	if  err := tsr.Ping(); err != nil {
		return "", err
	}

	return tsr.client.Get(key).Result()
}

<<<<<<< HEAD
func (tsr *typeString) GetRange(key string, start, end int64) (string,error) {
	if tsr == nil {
=======
func (tsr *TypeString) GetRange(key string, start, end int64) (string,error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 返回 key 中字符串值的子字符
	if  err := tsr.Ping(); err != nil {
		return "",err
	}
	return tsr.client.GetRange(key, start, end).Result()
}

func (tsr *TypeString) GetSet(key string, value interface{}) (string,error) {
	// 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
	if tsr == nil {
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return "",err
	}
	return tsr.client.GetSet(key, value).Result()
}


<<<<<<< HEAD
func (tsr *typeString) GetBit(key string, offset int64) (int64,error) {
	if tsr == nil {
=======
func (tsr *TypeString) GetBit(key string, offset int64) (int64,error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.client.GetBit(key, offset).Result()
}


<<<<<<< HEAD
func (tsr *typeString) SetBit(key string, offset int64, value int) (int64,error) {
	if tsr == nil {
=======
func (tsr *TypeString) SetBit(key string, offset int64, value int) (int64,error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.client.SetBit(key, offset, value).Result()
}
<<<<<<< HEAD
func (tsr *typeString) SetEx(key string, expiration time.Duration) (bool, error) {
	if tsr == nil {
=======
func (tsr *TypeString) SetEx(key string, expiration time.Duration) (bool, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.client.Expire(key,expiration).Result()
}
<<<<<<< HEAD
func (tsr *typeString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if tsr == nil {
=======
func (tsr *TypeString) SetNX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.client.SetNX(key , value ,expiration).Result()
}
<<<<<<< HEAD
func (tsr *typeString) MGet(keys ...string) ([]interface{},error) {
	if tsr == nil {
=======
func (tsr *TypeString) MGet(keys ...string) ([]interface{},error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return nil,err
	}
	return tsr.client.MGet(keys...).Result()
}
<<<<<<< HEAD
func (tsr *typeString) SetRange(key string, offset int64, value string) error {
	if tsr == nil {
=======
func (tsr *TypeString) SetRange(key string, offset int64, value string) error {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return err
	}
	return tsr.client.SetRange(key,offset,value).Err()
}
<<<<<<< HEAD
func (tsr *typeString) StrLen(key string) (int64,error) {
	if tsr == nil {
=======
func (tsr *TypeString) StrLen(key string) (int64,error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0,err
	}
	return tsr.client.StrLen(key).Result()
}
<<<<<<< HEAD
func (tsr *typeString) MSet(pairs ...interface{}) (string, error) {
	if tsr == nil {
=======
func (tsr *TypeString) MSet(pairs ...interface{}) (string, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return "" ,err
	}
	return tsr.client.MSet(pairs...).Result()
}
<<<<<<< HEAD
func (tsr *typeString) MSetNX(pairs ...interface{}) (bool, error) {
	if tsr == nil {
=======
func (tsr *TypeString) MSetNX(pairs ...interface{}) (bool, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return false, err
	}
	return tsr.client.MSetNX(pairs...).Result()
}
<<<<<<< HEAD
func (tsr *typeString) Incr(key string) (int64, error)  {
	if tsr == nil {
=======
func (tsr *TypeString) Incr(key string) (int64, error)  {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0 ,err
	}
	return tsr.client.Incr(key).Result()
}
<<<<<<< HEAD
func (tsr *typeString) IncrBy(key string, value int64) (int64, error) {
	if tsr == nil {
=======
func (tsr *TypeString) IncrBy(key string, value int64) (int64, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.IncrBy(key, value).Result()
}
<<<<<<< HEAD
func (tsr *typeString) IncrByFloat(key string, value float64) (float64, error) {
	if tsr == nil {
=======
func (tsr *TypeString) IncrByFloat(key string, value float64) (float64, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.IncrByFloat(key, value).Result()
}

<<<<<<< HEAD
func (tsr *typeString) Decr(key string) (int64, error) {
	if tsr == nil {
=======
func (tsr *TypeString) Decr(key string) (int64, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.Decr(key).Result()
}

<<<<<<< HEAD
func (tsr *typeString) DecrBy(key string, value int64) (int64, error)  {
	if tsr == nil {
=======
func (tsr *TypeString) DecrBy(key string, value int64) (int64, error)  {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.DecrBy(key, value).Result()
}

<<<<<<< HEAD
func (tsr *typeString) Append(key string, value string) (int64, error) {
	if tsr == nil {
=======
func (tsr *TypeString) Append(key string, value string) (int64, error) {
	if tsr.tsr == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tsr.Ping(); err != nil {
		return 0, err
	}
	return tsr.client.Append(key, value).Result()
}
