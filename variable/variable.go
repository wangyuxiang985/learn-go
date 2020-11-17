package main

import (
	"fmt"
	"math"
	"strings"
)

/*
Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用。
变量声明格式为：
1、	var 变量名 变量类型
2、 var(a string b int)
3、 var 变量名 = 变量值  //初始化，根据值推断类型
4、 var 变量名1，变量名2 = 值1，值2
5、 在函数内部 := 方式声明并初始化变量。

匿名变量（anonymous variable）: 用一个下划线_表示，匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
PS:
函数外的每个语句都必须以关键字开始（var、const、func等）
:=不能使用在函数外。
_多用于占位，表示忽略值。

常量：常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。
const pi = 3.1415
同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	pi1 = 100
	pi2
	pi3
)

 iota
iota是go语言的常量计数器，只能在常量的表达式中使用。 iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次
使用_跳过某些值
const (
	n1 = iota //0
	n2 = 100  // 100
	n3 = iota //2
	_
	n4        //4
)
*/

var name string
var age int
var (
	a string
	b int
	c bool
)

var flag  = false

const pi = 3.14
const (
	pi1 = 100
	pi2
	pi3 = 200
	pi4
)
const (
	n1 = iota
	n2 = 100
	n3 = iota
	_
	n4
)
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
)
const (
	d, e = iota + 1, iota +2 // 1, 2
	f, g                     // 2, 3
	h, z                     // 3, 4
)

func foo()(int, string) {
	return 10, "string"
}

func traversalString()  {
	s2 := "中国,china"
	for i2 := 0; i2 < len(s2); i2++ {
		fmt.Printf("%v(%c) ", s2[i2], s2[i2])
	}
	fmt.Println()

	/*range 遍历array，*array，string它返回两个值分别是数据的索引和值，遍历map时返回的两个值分别是key和value，遍历channel时，则只有一个返回数据*/
	for _, r := range s2 { //rune
		fmt.Printf("%v(%c) ", r, r)
	}

}

func main()  {
	//fmt.Println(a,b)
	flag := true
	fmt.Println(flag)

	i, _ := foo()
	_, s := foo()
	fmt.Println("i=", i)
	fmt.Println("s=", s)
	fmt.Println("pi4=", pi4)
	fmt.Println("n3=", n3)
	fmt.Println("n4=", n4)
	fmt.Println("KB=", KB)
	fmt.Println("MB=", MB)
	fmt.Println("GB=", GB)
	fmt.Println("d=", d)
	fmt.Println("e=", e)
	fmt.Println("f=", f)
	fmt.Println("g=", g)
	fmt.Println("h=", h)
	fmt.Println("z=", z)
	fmt.Println(math.MaxFloat32)

	s_ := `第一行
第二行
第三行`
	s_1 := "测试,test"
	fmt.Println(s_)
	fmt.Println("-----------------------")
	fmt.Println("字符串s_长度：", len(s_)) // 29 一个中文三个长度
	fmt.Println("字符串s_1长度：", len(s_1))
	sprintf := fmt.Sprintf("aaa%vccc", "中国") //拼接字符串
	fmt.Println("字符串拼接结果：", sprintf)

	split := strings.Split(s_1, ",") //分割
	fmt.Println(split[0])

	contains := strings.Contains(s_1, ",") // 判断是否包含
	fmt.Println("是否包含，：", contains)

	prefix := strings.HasPrefix(s_1, "测试") // 前缀判断
	fmt.Println("是否包含指定前缀：", prefix)

	index := strings.Index(s_1, ",")
	fmt.Println("s_1串中,出现的下标：", index)

	traversalString()
}
