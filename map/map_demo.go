package main

import (
	"fmt"
)

//map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用
/*
map语法：
 map[KeyType]ValueType
1.KeyType:表示键的类型。
2.ValueType:表示键对应的值的类型。
map类型的变量默认初始值为nil，需要使用make()函数来分配内存
make(map[KeyType]ValueType, [cap])
cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量
 */
//判断某个键是否存在  value, ok := map[key]
//若key存在则ok为true，value为对应的值；若key不存在，则ok为false，value为对应值类型的零值

//map的遍历
//使用for range遍历map
func main()  {
	fmt.Println("==========基本用法=========")
	basicUsage()
	fmt.Println("==============判断key是否存在===========")
	containsKey()
	fmt.Println("===========map遍历=============")
	rangeMap()

}

//遍历map
func rangeMap() {
	sourceMap := map[string]int{
		"李雷": 20,
		"丹尼": 50,
	}
	//遍历出来的顺序与添加顺序无关
	for k, v := range sourceMap {
		fmt.Println(k,v)
	}

	//只遍历key
	for k := range sourceMap {
		fmt.Println(k)
	}
	//只遍历value
	for _,v := range sourceMap {
		fmt.Println(v)
	}
}

//判断某个key是否存在
func containsKey() {
	sourceMap := map[string]int{
		"李雷": 20,
		"丹尼": 50,
	}
	value,ok :=sourceMap["李雷1"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println(value)
		fmt.Println("不存在")
	}
}
//map基本用法
func basicUsage() {
	sourcemMap := make(map[string]int, 8)
	sourcemMap["张三"] = 90
	sourcemMap["李四"] = 70
	fmt.Println(sourcemMap) //map[张三:90 李四:70]
	fmt.Println(sourcemMap["张三"])  //90
	fmt.Println(sourcemMap["小明"]) //0
	fmt.Printf("type of a:%T\n", sourcemMap) //type of a:map[string]int

	//声明时填充元素
	userInfo := map[string]string{
		"userName":"李明",
		"password":"1234",
	}
	fmt.Println(userInfo)
}
