package xredis

type TypeSet struct {
	ts *GoRedis
}


func (gr *GoRedis) NewSet() *TypeSet {
	return &TypeSet{gr	}
}


func (ts *TypeSet) SAdd(key string, members ...interface{}) (int64, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  0, err
	}

	return ts.ts.client.SAdd(key, members...).Result()
}

func (ts *TypeSet) SCard(key string) (int64, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  0, err
	}

	return ts.ts.client.SCard(key).Result()
}

func (ts *TypeSet) SDiff(key ...string) ([]string, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  nil, err
	}

	return ts.ts.client.SDiff(key...).Result()
}
func (ts *TypeSet) SDiffStore(destination string, keys ...string) (int64, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  0, err
	}

	return ts.ts.client.SDiffStore(destination, keys...).Result()
}

func (ts *TypeSet) SInter(keys ...string) ([]string, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  nil, err
	}

	return ts.ts.client.SInter(keys...).Result()
}

func (ts *TypeSet) SInterStore(destination string, keys ...string) (int64, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  0, err
	}

	return ts.ts.client.SInterStore(destination, keys...).Result()
}

func (ts *TypeSet) SIsMember(key string, member interface{}) (bool, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  false, err
	}

	return ts.ts.client.SIsMember(key, member).Result()
}


func (ts *TypeSet) SMembers(key string) ([]string, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  nil, err
	}

	return ts.ts.client.SMembers(key).Result()
}


func (ts *TypeSet) SMove(source, destination string, member interface{}) (bool, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  false, err
	}

	return ts.ts.client.SMove(source, destination,member).Result()
}


func (ts *TypeSet) SPop(key string) (string, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  "", err
	}

	return ts.ts.client.SPop(key).Result()
}


func (ts *TypeSet) SRandMember(key string) (string, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  "", err
	}

	return ts.ts.client.SRandMember(key).Result()
}


func (ts *TypeSet) SRandMemberN(key string, count int64) ([]string, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  nil, err
	}

	return ts.ts.client.SRandMemberN(key, count).Result()
}


func (ts *TypeSet) SRem(key string, members ...interface{}) (int64, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  0, err
	}

	return ts.ts.client.SRem(key, members...).Result()
}

func (ts *TypeSet) SUnion(keys ...string) ([]string, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  nil, err
	}

	return ts.ts.client.SUnion(keys...).Result()
}

func (ts *TypeSet) SUnionStore(destination string, keys ...string) (int64, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  0, err
	}

	return ts.ts.client.SUnionStore(destination, keys...).Result()
}


func (ts *TypeSet) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if ts.ts == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.ts.Ping(); err != nil {
		return  nil, 0, err
	}

	return ts.ts.client.SScan(key, cursor, match,count ).Result()
}









