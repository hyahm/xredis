package xredis

<<<<<<< HEAD
type typeSet struct {
	*GoRedis
=======
type TypeSet struct {
	ts *GoRedis
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
}


func (gr *GoRedis) NewSet() *TypeSet {
	return &TypeSet{gr	}
}


<<<<<<< HEAD
func (ts *typeSet) SAdd(key string, members ...interface{}) (int64, error) {
	if ts == nil {
=======
func (ts *TypeSet) SAdd(key string, members ...interface{}) (int64, error) {
	if ts.ts == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.Ping(); err != nil {
		return  0, err
	}

	return ts.client.SAdd(key, members...).Result()
}

<<<<<<< HEAD
func (ts *typeSet) SCard(key string) (int64, error) {
	if ts == nil {
=======
func (ts *TypeSet) SCard(key string) (int64, error) {
	if ts.ts == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.Ping(); err != nil {
		return  0, err
	}

	return ts.client.SCard(key).Result()
}

<<<<<<< HEAD
func (ts *typeSet) SDiff(key ...string) ([]string, error) {
	if ts == nil {
=======
func (ts *TypeSet) SDiff(key ...string) ([]string, error) {
	if ts.ts == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.Ping(); err != nil {
		return  nil, err
	}

	return ts.client.SDiff(key...).Result()
}
<<<<<<< HEAD
func (ts *typeSet) SDiffStore(destination string, keys ...string) (int64, error) {
	if ts == nil {
=======
func (ts *TypeSet) SDiffStore(destination string, keys ...string) (int64, error) {
	if ts.ts == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.Ping(); err != nil {
		return  0, err
	}

	return ts.client.SDiffStore(destination, keys...).Result()
}

<<<<<<< HEAD
func (ts *typeSet) SInter(keys ...string) ([]string, error) {
	if ts == nil {
=======
func (ts *TypeSet) SInter(keys ...string) ([]string, error) {
	if ts.ts == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.Ping(); err != nil {
		return  nil, err
	}

	return ts.client.SInter(keys...).Result()
}

<<<<<<< HEAD
func (ts *typeSet) SInterStore(destination string, keys ...string) (int64, error) {
	if ts == nil {
=======
func (ts *TypeSet) SInterStore(destination string, keys ...string) (int64, error) {
	if ts.ts == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := ts.Ping(); err != nil {
		return  0, err
	}

	return ts.client.SInterStore(destination, keys...).Result()
}

<<<<<<< HEAD
func (ts *typeSet) SIsMember(key string, member interface{}) (bool, error) {
	if ts == nil {
=======
func (ts *TypeSet) SIsMember(key string, member interface{}) (bool, error) {
	if ts.ts == nil {
>>>>>>> e4d6c3adf0ae90e4a60b4c67b2c3d1c706f2965d
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := dts.Ping(); err != nil {
		return  false, err
	}

	return ts.client.SIsMember(key, member).Result()
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









