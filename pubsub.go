package xredis

type TypePubSub struct {
	*GoRedis
}

func (gr *GoRedis) PubSub() *TypePubSub {
	return &TypePubSub{gr}
}

// func (gr *TypePubSub) PSubscribe(channels ...string) (int64, error) {
// 	return gr.client.PSubscribe(channels...)
// }
