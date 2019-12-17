package xredis

type typeSet struct {
	*GoRedis
}


func (gr *GoRedis) Set() *typeSet {
	return &typeSet{gr	}
}


func (gr *typeSet) SAdd(key string, members ...interface{}) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  0, err
	}

	return gr.client.SAdd(key, members...).Result()
}

func (gr *typeSet) SCard(key string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  0, err
	}

	return gr.client.SCard(key).Result()
}

func (gr *typeSet) SDiff(key ...string) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  nil, err
	}

	return gr.client.SDiff(key...).Result()
}
func (gr *typeSet) SDiffStore(destination string, keys ...string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  0, err
	}

	return gr.client.SDiffStore(destination, keys...).Result()
}

func (gr *typeSet) SInter(keys ...string) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  nil, err
	}

	return gr.client.SInter(keys...).Result()
}

func (gr *typeSet) SInterStore(destination string, keys ...string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  0, err
	}

	return gr.client.SInterStore(destination, keys...).Result()
}

func (gr *typeSet) SIsMember(key string, member interface{}) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  false, err
	}

	return gr.client.SIsMember(key, member).Result()
}


func (gr *typeSet) SMembers(key string) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  nil, err
	}

	return gr.client.SMembers(key).Result()
}


func (gr *typeSet) SMove(source, destination string, member interface{}) (bool, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  false, err
	}

	return gr.client.SMove(source, destination,member).Result()
}


func (gr *typeSet) SPop(key string) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  "", err
	}

	return gr.client.SPop(key).Result()
}


func (gr *typeSet) SRandMember(key string) (string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  "", err
	}

	return gr.client.SRandMember(key).Result()
}


func (gr *typeSet) SRandMemberN(key string, count int64) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  nil, err
	}

	return gr.client.SRandMemberN(key, count).Result()
}


func (gr *typeSet) SRem(key string, members ...interface{}) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  0, err
	}

	return gr.client.SRem(key, members...).Result()
}

func (gr *typeSet) SUnion(keys ...string) ([]string, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  nil, err
	}

	return gr.client.SUnion(keys...).Result()
}

func (gr *typeSet) SUnionStore(destination string, keys ...string) (int64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  0, err
	}

	return gr.client.SUnionStore(destination, keys...).Result()
}


func (gr *typeSet) SScan(key string, cursor uint64, match string, count int64) ([]string, uint64, error) {
	if gr == nil {
		panic("please conn first")
	}
	// 向集合添加一个或多个成员, 返回添加的成员数
	if  err := gr.ping(); err != nil {
		return  nil, 0, err
	}

	return gr.client.SScan(key, cursor, match,count ).Result()
}









