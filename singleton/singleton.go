package singleton

import (
	"fmt"
	"sync"
)

type object struct {
	name string
}

var once sync.Once
var obj *object //单例指针

//公开方法 外包调用
func Instance() *object {
	once.Do(getObj)
	return obj
}

func getObj() {
	if obj == nil {
		obj = new(object)
		//可以做其他初始化事件
	}
}

//单例测试
func (obj *object) Test() {
	fmt.Println(obj.name)
}
