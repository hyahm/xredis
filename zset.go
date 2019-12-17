package xredis

import (
	"github.com/go-redis/redis/v7"
)

type typeZset struct {
	*GoRedis
}


func (gr *GoRedis) Zset() *typeZset {
	return &typeZset{gr	}
}


func (gr *typeZset) ZAdd(key string, members ...*redis.Z) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZAdd(key, members...).Result()
}

func (gr *typeZset) ZCard(key string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZCard(key).Result()
}


func (gr *typeZset) ZCount(key, min, max string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZCount(key, min, max).Result()
}


func (gr *typeZset) ZIncrBy(key string, increment float64, member string) (float64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZIncrBy(key, increment, member).Result()
}


func (gr *typeZset) ZInterStore(destination string, store *redis.ZStore) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZInterStore(destination, store).Result()
}


func (gr *typeZset) ZLexCount(key, min, max string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZLexCount(key, min, max).Result()
}


func (gr *typeZset) ZRange(key string, start, stop int64) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.ZRange(key, start, stop).Result()
}


func (gr *typeZset) ZRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.ZRangeByLex(key, opt).Result()
}


func (gr *typeZset) ZRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.ZRangeByScore(key, opt).Result()
}


func (gr *typeZset) ZRangeByScoreWithScores(key string, opt *redis.ZRangeBy) ([]redis.Z, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.ZRangeByScoreWithScores(key, opt).Result()
}


func (gr *typeZset) ZRank(key, member string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZRank(key, member).Result()
}

func (gr *typeZset) ZRem(key string, members ...interface{}) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZRem(key, members...).Result()
}


func (gr *typeZset) ZRemRangeByLex(key, min, max string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZRemRangeByLex(key, min, max).Result()
}


func (gr *typeZset) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZRemRangeByRank(key, start, stop).Result()
}


func (gr *typeZset) ZRemRangeByScore(key, min, max string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZRemRangeByScore(key, min, max).Result()
}


func (gr *typeZset) ZRevRange(key string, start, stop int64) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.ZRevRange(key, start, stop).Result()
}


func (gr *typeZset) ZRevRangeByLex(key string, opt *redis.ZRangeBy) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.ZRevRangeByLex(key, opt).Result()
}


func (gr *typeZset) ZRevRangeByScore(key string, opt *redis.ZRangeBy) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, err
	}

	return gr.client.ZRevRangeByScore(key, opt).Result()
}


func (gr *typeZset) ZScore(key, member string) (float64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZScore(key, member).Result()
}


func (gr *typeZset) ZUnionStore(dest string, store *redis.ZStore) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return 0, err
	}

	return gr.client.ZUnionStore(dest, store).Result()
}


func (gr *typeZset) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if gr == nil {
		panic("please conn first")
	}
	if  err := gr.Ping(); err != nil {
		return nil, 0, err
	}

	return gr.client.ZScan(key, cursor, match,count).Result()
}
