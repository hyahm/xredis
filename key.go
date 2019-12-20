package xredis

import "time"

type TypeKey struct {
	tk *GoRedis
}


func (gr *GoRedis) NewKey() *TypeKey {
	return &TypeKey{gr	}
}

func (tk *TypeKey) Del(keys ...string) (int64, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  0, err
	}
	return tk.tk.client.Del(keys...).Result()
}


func (tk *TypeKey) Dump(key string) (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.Dump(key).Result()
}


func (tk *TypeKey) RPushX(key string, values ...interface{}) (int64, error) {
	// 为已存在的列表添加值
	if tk.tk == nil {
		panic("please conn first")
	}
	if err := tk.tk.Ping(); err != nil {
		return 0, err
	}

	return tk.tk.client.RPushX(key, values...).Result()
}


func (tk *TypeKey) Expire(key string, expiration time.Duration) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.Expire(key, expiration).Result()
}



func (tk *TypeKey) ExpireAt(key string, tm time.Time) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.ExpireAt(key, tm).Result()
}



func (tk *TypeKey) Keys(pattern string) ([]string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  nil, err
	}
	return tk.tk.client.Keys(pattern).Result()
}


func (tk *TypeKey) Move(key string, db int) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.Move(key,db).Result()
}


func (tk *TypeKey) Persist(key string) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.Persist(key).Result()
}

func (tk *TypeKey) TTL(key string) (time.Duration, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  0, err
	}
	return tk.tk.client.TTL(key).Result()
}

func (tk *TypeKey) RandomKey() (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.RandomKey().Result()
}


func (tk *TypeKey) Rename(key, newkey string) (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.Rename(key,newkey).Result()
}


func (tk *TypeKey) RenameNX(key, newkey string) (bool, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  false, err
	}
	return tk.tk.client.RenameNX(key,newkey).Result()
}


func (tk *TypeKey) Type(key string) (string, error) {
	if tk.tk == nil {
		panic("please conn first")
	}
	if  err := tk.tk.Ping(); err != nil {
		return  "", err
	}
	return tk.tk.client.Type(key).Result()
}


