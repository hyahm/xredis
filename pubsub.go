package xredis
type typePubSub struct {
	*GoRedis
}


func (gr *GoRedis) PubSub() *typePubSub {
	return &typePubSub{gr	}
}

//func (gr *typeSet) PSubscribe(channels ...string) (int64, error) {
//	if gr == nil {
//		panic("please conn first")
//	}
//	// 向集合添加一个或多个成员, 返回添加的成员数
//	if  err := gr.Ping(); err != nil {
//		return  0, err
//	}
//
//	return gr.client.PSubscribe(channels...)
//}
