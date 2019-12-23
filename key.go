package xredis

import "time"

<<<<<<< HEAD

func (tk *GoRedis) Del(keys ...string) (int64, error) {
	if tk == nil {
=======
type TypeKey struct {
	tk *GoRedis
}


func (gr *GoRedis) NewKey() *TypeKey {
	return &TypeKey{gr	}
}

func (tk *TypeKey) Del(keys ...string) (int64, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  0, err
	}
	return tk.client.Del(keys...).Result()
}


<<<<<<< HEAD
func (tk *GoRedis) Dump(key string) (string, error) {
	if tk == nil {
=======
func (tk *TypeKey) Dump(key string) (string, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  "", err
	}
	return tk.client.Dump(key).Result()
}


<<<<<<< HEAD
func (tk *GoRedis) RPushX(key string, values ...interface{}) (int64, error) {
=======
func (tk *TypeKey) RPushX(key string, values ...interface{}) (int64, error) {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
	// 为已存在的列表添加值
	if tk == nil {
		panic("please conn first")
	}
	if err := tk.Ping(); err != nil {
		return 0, err
	}

	return tk.client.RPushX(key, values...).Result()
}


<<<<<<< HEAD
func (tk *GoRedis) Expire(key string, expiration time.Duration) (bool, error) {
	if tk == nil {
=======
func (tk *TypeKey) Expire(key string, expiration time.Duration) (bool, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  false, err
	}
	return tk.client.Expire(key, expiration).Result()
}



<<<<<<< HEAD
func (tk *typeKey) ExpireAt(key string, tm time.Time) (bool, error) {
	if tk == nil {
=======
func (tk *TypeKey) ExpireAt(key string, tm time.Time) (bool, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  false, err
	}
	return tk.client.ExpireAt(key, tm).Result()
}



<<<<<<< HEAD
func (tk *typeKey) Keys(pattern string) ([]string, error) {
	if tk == nil {
=======
func (tk *TypeKey) Keys(pattern string) ([]string, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  nil, err
	}
	return tk.client.Keys(pattern).Result()
}


<<<<<<< HEAD
func (tk *typeKey) Move(key string, db int) (bool, error) {
	if tk == nil {
=======
func (tk *TypeKey) Move(key string, db int) (bool, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  false, err
	}
	return tk.client.Move(key,db).Result()
}


<<<<<<< HEAD
func (tk *typeKey) Persist(key string) (bool, error) {
	if tk == nil {
=======
func (tk *TypeKey) Persist(key string) (bool, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  false, err
	}
	return tk.client.Persist(key).Result()
}

<<<<<<< HEAD
func (tk *typeKey) TTL(key string) (time.Duration, error) {
	if tk == nil {
=======
func (tk *TypeKey) TTL(key string) (time.Duration, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  0, err
	}
	return tk.client.TTL(key).Result()
}

<<<<<<< HEAD
func (tk *typeKey) RandomKey() (string, error) {
	if tk == nil {
=======
func (tk *TypeKey) RandomKey() (string, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  "", err
	}
	return tk.client.RandomKey().Result()
}


<<<<<<< HEAD
func (tk *typeKey) Rename(key, newkey string) (string, error) {
	if tk == nil {
=======
func (tk *TypeKey) Rename(key, newkey string) (string, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  "", err
	}
	return tk.client.Rename(key,newkey).Result()
}


<<<<<<< HEAD
func (tk *typeKey) RenameNX(key, newkey string) (bool, error) {
	if tk == nil {
=======
func (tk *TypeKey) RenameNX(key, newkey string) (bool, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  false, err
	}
	return tk.client.RenameNX(key,newkey).Result()
}


<<<<<<< HEAD
func (tk *typeKey) Type(key string) (string, error) {
	if tk == nil {
=======
func (tk *TypeKey) Type(key string) (string, error) {
	if tk.tk == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	if  err := tk.Ping(); err != nil {
		return  "", err
	}
	return tk.client.Type(key).Result()
}


