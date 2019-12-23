package xredis

type TypeHash struct {
	th *GoRedis
}


func (gr *GoRedis) NewHash() *TypeHash {
	return &TypeHash{gr	}
}

func (th *TypeHash) HDel(key string, fields ...string) error {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return  err
	}

	return th.th.client.HDel(key, fields...).Err()
}
func (th *TypeHash) HExists(key, field string) (bool, error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return false, err
	}

	return th.th.client.HExists(key, field).Result()
}

func (th *TypeHash) HGet(key, field string) (string,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return "", err
	}

	return th.th.client.HGet(key, field).Result()
}

func (th *TypeHash) HGetAll(key string) (map[string]string,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return nil, err
	}

	return th.th.client.HGetAll(key).Result()
}
func (th *TypeHash) HIncrBy(key, field string, incr int64) (int64,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return 0, err
	}

	return th.th.client.HIncrBy(key, field, incr).Result()
}
func (th *TypeHash) HIncrByFloat(key, field string, incr float64) (float64,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return 0, err
	}

	return th.th.client.HIncrByFloat(key, field, incr).Result()
}
func (th *TypeHash) HKeys(key string) ([]string,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return nil, err
	}

	return th.th.client.HKeys(key).Result()
}
func (th *TypeHash) HLen(key string) (int64,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return 0, err
	}

	return th.th.client.HLen(key).Result()
}
func (th *TypeHash) HMGet(key string,  fields ...string) ([]interface{},error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return nil, err
	}

	return th.th.client.HMGet(key, fields...).Result()
}

func (th *TypeHash) HMSet(key string, fields map[string]interface{}) (string,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return "", err
	}

	return th.th.client.HMSet(key, fields).Result()
}
func (th *TypeHash) HSet(key, field string, value interface{}) (bool,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return false, err
	}

	return th.th.client.HSet(key, field, value).Result()
}

func (th *TypeHash) HSetNX(key, field string, value interface{}) (bool,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return false, err
	}

	return th.th.client.HSetNX(key, field, value).Result()
}

func (th *TypeHash) HVals(key string) ([]string,error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return nil, err
	}

	return th.th.client.HVals(key).Result()
}

func (th *TypeHash) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if th.th == nil {
		panic("please conn first")
	}
	if  err := th.th.Ping(); err != nil {
		return nil, 0, err
	}


	return th.th.client.HScan(key,cursor,match,count).Result()
}




