package xredis

type typeTable struct {
	*GoRedis
}


func (gr *GoRedis) Table() *typeTable {
	return &typeTable{gr}
}

// 多key共享一个值
