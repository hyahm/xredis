package xredis

type TypeHash struct {
	*GoRedis
}

func (gr *GoRedis) NewHash() *TypeHash {
	return &TypeHash{gr}
}

func (th *TypeHash) HDel(key string, fields ...string) error {
	return th.client.HDel(key, fields...).Err()
}
func (th *TypeHash) HExists(key, field string) (bool, error) {
	return th.client.HExists(key, field).Result()
}

func (th *TypeHash) HGet(key, field string) (string, error) {
	return th.client.HGet(key, field).Result()
}

func (th *TypeHash) HGetAll(key string) (map[string]string, error) {
	return th.client.HGetAll(key).Result()
}
func (th *TypeHash) HIncrBy(key, field string, incr int64) (int64, error) {
	return th.client.HIncrBy(key, field, incr).Result()
}
func (th *TypeHash) HIncrByFloat(key, field string, incr float64) (float64, error) {
	return th.client.HIncrByFloat(key, field, incr).Result()
}
func (th *TypeHash) HKeys(key string) ([]string, error) {
	return th.client.HKeys(key).Result()
}
func (th *TypeHash) HLen(key string) (int64, error) {
	return th.client.HLen(key).Result()
}
func (th *TypeHash) HMGet(key string, fields ...string) ([]interface{}, error) {
	return th.client.HMGet(key, fields...).Result()
}

func (th *TypeHash) HMSet(key string, fields map[string]interface{}) (string, error) {
	return th.client.HMSet(key, fields).Result()
}
func (th *TypeHash) HSet(key, field string, value interface{}) (bool, error) {
	return th.client.HSet(key, field, value).Result()
}

func (th *TypeHash) HSetNX(key, field string, value interface{}) (bool, error) {
	return th.client.HSetNX(key, field, value).Result()
}

func (th *TypeHash) HVals(key string) ([]string, error) {
	return th.client.HVals(key).Result()
}

func (th *TypeHash) HScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return th.client.HScan(key, cursor, match, count).Result()
}
