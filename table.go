package xredis

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type typeTable struct {
	tt *GoRedis
	prefix string // 前缀
	key *TypeSet
	kt *TypeString
	mu sync.RWMutex
}

// 字段前缀默认table, 无法修改
func (gr *GoRedis) NewTable() *typeTable {

	return &typeTable{
		tt: gr,
		prefix: "table",
		kt: gr.NewStr(),
		key: gr.NewSet(),
		mu: sync.RWMutex{},
	}
}

// 插入表结构， 只支持基本数据类型， 这个数据是可以标记key的
func (tt *typeTable) InsertTemplate(name string, table interface{}) error {
	t := reflect.ValueOf(table)
	if t.Kind() != reflect.Struct {
		return  errors.New("must be struct")
	}

	// 插入redis
	kt := tt.tt.NewStr()
	value, err := kt.Set(tt.prefix + name, t.Type().String(), 0)
	if err != nil {
		return err
	}

	fmt.Println("set: ",value)
	return nil
}

func (tt *typeTable) SetPrefix(name string, prefix string) error {
	// 只有没有值
	_, err := tt.kt.Get(tt.prefix + name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (tt *typeTable) Key(name string, keys ...interface{}) error {
	// 设置key
	// 只有没有值
	_, err := tt.kt.Get(tt.prefix + name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//if _, ok := reflect.TypeOf(tt.table[name]).FieldByName(key); ok {
	//	tt.key = append(tt.key, key)
	//	// 增加key， 放入哪个类型中
	//}

	return nil
}

func (tt *typeTable) GetKeys(name string) ([]string, error) {

	// 只有没有值
	_, err := tt.kt.Get(tt.prefix + name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	values, err := tt.key.SMembers(name)
	if err != nil {
		return nil, err
	}
	return values, nil
}

func (tt *typeTable) Add(name string, data interface{}) error {
	// 添加数据并创建key
	// value
	//
	// 先判断结构是否一样
	fmt.Println(reflect.TypeOf(data).Kind())
	fmt.Println(reflect.TypeOf(data).Kind())
	return nil
}

// 获取值
func (tt *typeTable) Get(name string, key string, value interface{}) error {
	// 添加数据并创建key
	// value
	//
	return nil
}

// 设置值
func (tt *typeTable) Set(name string, data interface{}, oldvalue interface{}, value interface{}) error {
	// 添加数据并创建key
	// value
	//
	return nil
}
