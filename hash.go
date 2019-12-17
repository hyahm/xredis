package xredis

type typeHash struct {
	*GoRedis
}


func (gr *GoRedis) Hash() *typeHash {
	return &typeHash{gr	}
}

func (gr *typeHash) HDel(key string, fields ...string) error {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return  err
	}

	return gr.client.HDel(key, fields...).Err()
}
func (gr *typeHash) HExists(key, field string) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return false, err
	}

	return gr.client.HExists(key, field).Result()
}

func (gr *typeHash) HGet(key, field string) (string,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.HGet(key, field).Result()
}

func (gr *typeHash) HGetAll(key string) (map[string]string,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.HGetAll(key).Result()
}
func (gr *typeHash) HIncrBy(key, field string, incr int64) (int64,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.HIncrBy(key, field, incr).Result()
}
func (gr *typeHash) HIncrByFloat(key, field string, incr float64) (float64,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.HIncrByFloat(key, field, incr).Result()
}
func (gr *typeHash) HKeys(key string) ([]string,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.HKeys(key).Result()
}
func (gr *typeHash) HLen(key string) (int64,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.HLen(key).Result()
}
func (gr *typeHash) HMGet(key string,  fields ...string) ([]interface{},error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.HMGet(key, fields...).Result()
}

func (gr *typeHash) HMSet(key string, fields map[string]interface{}) (string,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return "", err
	}

	return gr.client.HMSet(key, fields).Result()
}
func (gr *typeHash) HSet(key, field string, value interface{}) (bool,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return false, err
	}

	return gr.client.HSet(key, field, value).Result()
}

func (gr *typeHash) HSetNX(key, field string, value interface{}) (bool,error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return false, err
	}

	return gr.client.HSetNX(key, field, value).Result()
}

func (gr *typeHash) HVals(key string) ([]string,error) {
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.HVals(key).Result()
}

func (gr *typeHash) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil,0,  err
	}

	return gr.client.HScan(key,cursor,match,count).Result()
}




