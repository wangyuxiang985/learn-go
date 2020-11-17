package main

import "fmt"

//一维数组
/*
---------初始化-----------------
全局：
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}
局部：
a := [3]int{1, 2}           // 未初始化元素值为 0。
b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。
d := [...]struct {
	name string
	age  uint8
}{
	{"user1", 10}, // 可省略元素类型。
	{"user2", 20}, // 别忘了最后一行的逗号。
}

*/
var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3}
var str = [5]string{2: "hello world", 3: "tony"}

func main()  {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3}
	c := [5]int{2: 100, 3:100}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"zhangsan", 10},
		{"李四", 20},
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
	fmt.Println(len(a), len(b))
	fmt.Println(str[2])

}
