package main

import (
	"fmt"
)

//在 Go 中，与 C 数组变量隐式作为指针使用不同，Go 数组是值类型，赋值和函数传参操作都会复制整个数组数据
//Slice 的数据结构定义如下：
/**
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
 */
//切片的结构体由3部分构成，Pointer 是指向一个数组的指针，len 代表当前切片的长度，cap 是当前切片的容量。cap 总是大于等于 len 的

//用字面量创建的一个 len = 6，cap = 6 的切片 slice := []int{1,2,3,4,5,6}
//需要注意的是 [ ] 里面不要写数组的容量，因为如果写了个数以后就是数组了，而不是切片了
//空切片和 nil 切片的区别在于，空切片指向的地址不是nil，指向的是一个内存地址，但是它没有分配任何内存空间，即底层元素包含0个元素
//不管是使用 nil 切片还是空切片，对其调用内置函数 append，len 和 cap 的效果都是一样的

func main()  {

	arrayA := [2]int{100, 200}
	var arrayB [2]int

	arrayB = arrayA

	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA) //arrayA : 0xc00000a0d0 , [100 200]
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB) //arrayB : 0xc00000a0e0 , [100 200]

	testArray(arrayA) //testArray/ints : 0xc00000a130 , [100 200]
	//上述三个内存地址都不同，这也就验证了 Go 中数组赋值和函数传参都是值复制的。


	arrayC := []int{1, 2}
	testArrayPoint(&arrayC) //传指针 func Array : 0xc0000044c0 , [1 2]
	arrayD := arrayC[:]
	testArrayPoint(&arrayD) //传切片 func Array : 0xc000004500 , [1 102]
	fmt.Printf("arrayC : %p , %v\n", &arrayC, arrayC) // arrayC : 0xc0000044c0 , [1 202]

}

func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}

func testArray(ints [2]int) {
	fmt.Printf("testArray/ints : %p , %v\n", &ints, ints)
}
