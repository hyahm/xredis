package xredis

import "time"

type typeKey struct {
	*GoRedis
}


func (gr *GoRedis) Key() *typeKey {
	return &typeKey{gr	}
}

func (gr *typeKey) Del(keys ...string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  0, err
	}
	return gr.client.Del(keys...).Result()
}


func (gr *typeKey) Dump(key string) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  "", err
	}
	return gr.client.Dump(key).Result()
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


func (gr *typeKey) Expire(key string, expiration time.Duration) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  false, err
	}
	return gr.client.Expire(key, expiration).Result()
}



func (gr *typeKey) ExpireAt(key string, tm time.Time) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  false, err
	}
	return gr.client.ExpireAt(key, tm).Result()
}



func (gr *typeKey) Keys(pattern string) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  nil, err
	}
	return gr.client.Keys(pattern).Result()
}


func (gr *typeKey) Move(key string, db int) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  false, err
	}
	return gr.client.Move(key,db).Result()
}


func (gr *typeKey) Persist(key string) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  false, err
	}
	return gr.client.Persist(key).Result()
}

func (gr *typeKey) TTL(key string) (time.Duration, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  0, err
	}
	return gr.client.TTL(key).Result()
}

func (gr *typeKey) RandomKey() (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  "", err
	}
	return gr.client.RandomKey().Result()
}


func (gr *typeKey) Rename(key, newkey string) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  "", err
	}
	return gr.client.Rename(key,newkey).Result()
}


func (gr *typeKey) RenameNX(key, newkey string) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  false, err
	}
	return gr.client.RenameNX(key,newkey).Result()
}


func (gr *typeKey) Type(key string) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  "", err
	}
	return gr.client.Type(key).Result()
}


