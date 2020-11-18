package main

import "fmt"

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
func main()  {
	fmt.Println("==========基本用法=========")
	basicUsage()

}

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
