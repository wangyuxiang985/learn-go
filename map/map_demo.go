package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
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

//使用delete()函数删除键值对
/*
delete()函数的格式如下：
delete(map, key)
map:表示要删除键值对的map
key:表示要删除的键值对的键
 */
func main()  {
	fmt.Println("==========基本用法=========")
	basicUsage()
	fmt.Println("==============判断key是否存在===========")
	containsKey()
	fmt.Println("===========map遍历=============")
	rangeMap()
	fmt.Println("==========delete()删除键值对")
	deleteMapKey()
	fmt.Println("==========按照指定顺序遍历map==========")
	sortedRangeMap()

}

//按照指定顺序遍历map
func sortedRangeMap() {
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	var sourceMap = make(map[string]int, 200)

	for i := 0;i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		sourceMap[key] = value
	}
	//将map中的key取出存入数组并排序
	var keys = make([]string, 0, 200)
	for key := range sourceMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, sourceMap[key])
	}
}

//使用delete()函数删除键值对
func deleteMapKey() {
	sourceMap := map[string]int{
		"李雷": 20,
		"丹尼": 50,
	}
	delete(sourceMap,"李雷")
	delete(sourceMap,"李雷1")
	fmt.Println(sourceMap)
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
