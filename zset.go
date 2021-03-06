package xredis

type TypeZset struct {
	*GoRedis
}

// func (gr *GoRedis) NewZset() *TypeZset {
// 	return &TypeZset{gr}
// }

// func (tz *TypeZset) ZAdd(key string, members ...redis.Z) (int64, error) {
// 	return tz.client.ZAdd(key, members...).Result()
// }

// func (tz *TypeZset) ZCard(key string) (int64, error) {
// 	return tz.client.ZCard(key).Result()
// }

// func (tz *TypeZset) ZCount(key, min, max string) (int64, error) {
// 	return tz.client.ZCount(key, min, max).Result()
// }

// func (tz *TypeZset) ZIncrBy(key string, increment float64, member string) (float64, error) {
// 	return tz.client.ZIncrBy(key, increment, member).Result()
// }

// func (tz *TypeZset) ZInterStore(destination string, store redis.ZStore, keys ...string) (int64, error) {
// 	return tz.client.ZInterStore(destination, store, keys...).Result()
// }

// func (tz *TypeZset) ZLexCount(key, min, max string) (int64, error) {
// 	return tz.client.ZLexCount(key, min, max).Result()
// }

// func (tz *TypeZset) ZRange(key string, start, stop int64) ([]string, error) {
// 	return tz.client.ZRange(key, start, stop).Result()
// }

// func (tz *TypeZset) ZRangeByLex(key string, opt redis.ZRangeBy) ([]string, error) {
// 	return tz.client.ZRangeByLex(key, opt).Result()
// }

// func (tz *TypeZset) ZRangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
// 	return tz.client.ZRangeByScore(key, opt).Result()
// }

// func (tz *TypeZset) ZRangeByScoreWithScores(key string, opt redis.ZRangeBy) ([]redis.Z, error) {
// 	return tz.client.ZRangeByScoreWithScores(key, opt).Result()
// }

// func (tz *TypeZset) ZRank(key, member string) (int64, error) {
// 	return tz.client.ZRank(key, member).Result()
// }

// func (tz *TypeZset) ZRem(key string, members ...interface{}) (int64, error) {
// 	return tz.client.ZRem(key, members...).Result()
// }

// func (tz *TypeZset) ZRemRangeByLex(key, min, max string) (int64, error) {
// 	return tz.client.ZRemRangeByLex(key, min, max).Result()
// }

// func (tz *TypeZset) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
// 	return tz.client.ZRemRangeByRank(key, start, stop).Result()
// }

// func (tz *TypeZset) ZRemRangeByScore(key, min, max string) (int64, error) {
// 	return tz.client.ZRemRangeByScore(key, min, max).Result()
// }

// func (tz *TypeZset) ZRevRange(key string, start, stop int64) ([]string, error) {
// 	return tz.client.ZRevRange(key, start, stop).Result()
// }

// func (tz *TypeZset) ZRevRangeByLex(key string, opt redis.ZRangeBy) ([]string, error) {
// 	return tz.client.ZRevRangeByLex(key, opt).Result()
// }

// func (tz *TypeZset) ZRevRangeByScore(key string, opt redis.ZRangeBy) ([]string, error) {
// 	return tz.client.ZRevRangeByScore(key, opt).Result()
// }

// func (tz *TypeZset) ZScore(key, member string) (float64, error) {
// 	return tz.client.ZScore(key, member).Result()
// }

// func (tz *TypeZset) ZUnionStore(dest string, store redis.ZStore, keys ...string) (int64, error) {
// 	return tz.client.ZUnionStore(dest, store, keys...).Result()
// }

// func (tz *TypeZset) ZScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
// 	return tz.client.ZScan(key, cursor, match, count).Result()
// }
