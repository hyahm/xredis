package xredis

import "time"

type typeKey struct {
	tk *GoRedis
}


func (gr *GoRedis) NewKey() *typeKey {
	return &typeKey{gr	}
}

func (tk *typeKey) Del(keys ...string) (int64, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  0, err
	}
	return tk.tk.client.Del(keys...).Result()
}


func (tk *typeKey) Dump(key string) (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.Dump(key).Result()
}


func (tk *typeKey) RPushX(key string, values ...interface{}) (int64, error) {
	// 为已存在的列表添加值
	if tk.tk == nil {
		panic("please conn first")
	}
	if err := tk.tk.Ping(); err != nil {
		return 0, err
	}

	return tk.tk.client.RPushX(key, values...).Result()
}


func (tk *typeKey) Expire(key string, expiration time.Duration) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.Expire(key, expiration).Result()
}



func (tk *typeKey) ExpireAt(key string, tm time.Time) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.ExpireAt(key, tm).Result()
}



func (tk *typeKey) Keys(pattern string) ([]string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  nil, err
	}
	return tk.tk.client.Keys(pattern).Result()
}


func (tk *typeKey) Move(key string, db int) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.Move(key,db).Result()
}


func (tk *typeKey) Persist(key string) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.Persist(key).Result()
}

func (tk *typeKey) TTL(key string) (time.Duration, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  0, err
	}
	return tk.tk.client.TTL(key).Result()
}

func (tk *typeKey) RandomKey() (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.RandomKey().Result()
}


func (tk *typeKey) Rename(key, newkey string) (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.Rename(key,newkey).Result()
}


func (tk *typeKey) RenameNX(key, newkey string) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.RenameNX(key,newkey).Result()
}


func (tk *typeKey) Type(key string) (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.Type(key).Result()
}


