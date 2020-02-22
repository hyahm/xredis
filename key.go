package xredis

import "time"

func (tk *GoRedis) Del(keys ...string) (int64, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return 0, err
	}
	return tk.client.Del(keys...).Result()
}

func (tk *GoRedis) Dump(key string) (string, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return "", err
	}
	return tk.client.Dump(key).Result()
}

func (tk *GoRedis) RPushX(key string, value interface{}) (int64, error) {
	// 为已存在的列表添加值
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return 0, err
	}

	return tk.client.RPushX(key, value).Result()
}

func (tk *GoRedis) Expire(key string, expiration time.Duration) (bool, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return false, err
	}
	return tk.client.Expire(key, expiration).Result()
}

func (tk *GoRedis) ExpireAt(key string, tm time.Time) (bool, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return false, err
	}
	return tk.client.ExpireAt(key, tm).Result()
}

func (tk *GoRedis) Keys(pattern string) ([]string, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return nil, err
	}
	return tk.client.Keys(pattern).Result()
}

func (tk *GoRedis) Move(key string, db int) (bool, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return false, err
	}
	return tk.client.Move(key, db).Result()
}

func (tk *GoRedis) Persist(key string) (bool, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return false, err
	}
	return tk.client.Persist(key).Result()
}

func (tk *GoRedis) TTL(key string) (time.Duration, error) {
	if tk == nil {

		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return 0, err
	}
	return tk.client.TTL(key).Result()
}

func (tk *GoRedis) RandomKey() (string, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return "", err
	}
	return tk.client.RandomKey().Result()
}

func (tk *GoRedis) Rename(key, newkey string) (string, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return "", err
	}
	return tk.client.Rename(key, newkey).Result()
}

func (tk *GoRedis) RenameNX(key, newkey string) (bool, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return false, err
	}
	return tk.client.RenameNX(key, newkey).Result()
}

func (tk *GoRedis) Type(key string) (string, error) {
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return "", err
	}
	return tk.client.Type(key).Result()
}
