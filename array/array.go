package main

import (
	"fmt"
)

//1. 数组：是同一种数据类型的固定长度的序列。
//2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
//3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
//4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
//for i := 0; i < len(a); i++ {}
//for index, v := range a {}
//5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
//6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
//7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
//8.指针数组 [n]*T，数组指针 *[n]T。

func main() {
	//测试值交换 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
	fmt.Println("-----------测试值交换---------")
	test1 := [2]int{1, 2}
	fmt.Printf("test1: %p\n", &test1)
	test(test1)
	fmt.Println(test1)
	fmt.Println("===========测试指针数组交换拷贝")
	test2 := [3]int{1, 2}
	fmt.Println(&test2)
	printArr(&test2) //& 是取地址符号 , 即取得某个变量的地址 , 如 ; &a
	fmt.Println(test2)
}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
	fmt.Println(x)
}

func printArr(array *[3]int) {
	fmt.Println(&array)
	array[0] = 1000
	for i,v := range array {
		fmt.Println(i, v)
	}
}
