package xredis

import (
	"github.com/go-redis/redis/v7"
)

type typeZset struct {
	tz *GoRedis
}


func (gr *GoRedis) NewZset() *typeZset {
	return &typeZset{gr	}
}


func (tz *typeZset) ZAdd(key string, members ...*redis.Z) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZAdd(key, members...).Result()
}

func (tz *typeZset) ZCard(key string) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZCard(key).Result()
}


func (tz *typeZset) ZCount(key, min, max string) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZCount(key, min, max).Result()
}


func (tz *typeZset) ZIncrBy(key string, increment float64, member string) (float64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZIncrBy(key, increment, member).Result()
}


func (tz *typeZset) ZInterStore(destination string, store *redis.ZStore) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZInterStore(destination, store).Result()
}


func (tz *typeZset) ZLexCount(key, min, max string) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZLexCount(key, min, max).Result()
}


func (tz *typeZset) ZRange(key string, start, stop int64) ([]string, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, err
	}

	return tz.tz.client.ZRange(key, start, stop).Result()
}


func (tz *typeZset) ZRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, err
	}

	return tz.tz.client.ZRangeByLex(key, opt).Result()
}


func (tz *typeZset) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, err
	}

	return tz.tz.client.ZRangeByScore(key, opt).Result()
}


func (tz *typeZset) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, err
	}

	return tz.tz.client.ZRangeByScoreWithScores(key, opt).Result()
}


func (tz *typeZset) ZRank(key, member string) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZRank(key, member).Result()
}

func (tz *typeZset) ZRem(key string, members ...interface{}) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZRem(key, members...).Result()
}


func (tz *typeZset) ZRemRangeByLex(key, min, max string) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZRemRangeByLex(key, min, max).Result()
}


func (tz *typeZset) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZRemRangeByRank(key, start, stop).Result()
}


func (tz *typeZset) ZRemRangeByScore(key, min, max string) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZRemRangeByScore(key, min, max).Result()
}


func (tz *typeZset) ZRevRange(key string, start, stop int64) ([]string, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, err
	}

	return tz.tz.client.ZRevRange(key, start, stop).Result()
}


func (tz *typeZset) ZRevRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, err
	}

	return tz.tz.client.ZRevRangeByLex(key, opt).Result()
}


func (tz *typeZset) ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, err
	}

	return tz.tz.client.ZRevRangeByScore(key, opt).Result()
}


func (tz *typeZset) ZScore(key, member string) (float64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZScore(key, member).Result()
}


func (tz *typeZset) ZUnionStore(dest string, store *redis.ZStore) (int64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return 0, err
	}

	return tz.tz.client.ZUnionStore(dest, store).Result()
}


func (tz *typeZset) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if tz.tz == nil {
		panic("please conn first")
	}
	if  err := tz.tz.Ping(); err != nil {
		return nil, 0, err
	}

	return tz.tz.client.ZScan(key, cursor, match,count).Result()
}
