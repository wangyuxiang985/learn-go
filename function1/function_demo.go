package main

import "fmt"

/*
golang函数特点:
    • 无需声明原型。
    • 支持不定 变参。
    • 支持多返回值。
    • 支持命名返回参数。
    • 支持匿名函数和闭包。
    • 函数也是一种类型，一个函数可以赋值给变量。

    • 不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
    • 不支持 重载 (overload)
    • 不支持 默认参数 (default parameter)。

函数声明：
函数声明包含一个函数名，参数列表， 返回值列表和函数体。如果函数没有返回值，则返回列表可以省略。函数从第一条语句开始执行，直到执行return语句或者执行函数的最后一条语句。
函数可以没有参数或接受多个参数。注意类型在变量名之后 。两个或多个连续的函数命名参数是同一类型，则除了最后一个类型之外，其他都可以省略。函数可以返回任意数量的返回值。
使用关键字 func 定义函数，左大括号依旧不能另起一行。
	func test(x, y int, s string) (int, string) {
		// 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
		n := x + y
		return n, fmt.Sprintf(s, n)
	}
函数是第一类对象，可作为参数传递。建议将复杂签名定义为函数类型，以便于阅读。
有返回值的函数，必须有明确的终止语句，否则会引发编译错误
没有函数体的函数声明表示该函数不是以Go实现的。这样的声明定义了函数标识符
例如：
	package math
	func Sin(x float64) float //implemented in assembly language

参数：
函数可以通过两种方式来传递参数：
	值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
	引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
PS：
	无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝
	map、slice、chan、指针、interface默认以引用的方式传递。
不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个，在参数后加上“…”即可。
func myfunc(args ...int) {    //0个或多个参数} //其中args是一个slice，我们可以通过arg[index]依次访问所有参数,通过len(arg)来判断传递参数的个数.
任意类型的不定参数： 就是函数的参数和每个参数的类型都不是固定的。
	用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的。func myfunc(args ...interface{}) {}
使用 slice 对象做变参时，必须展开。（slice...）
*/
func test(fn func() int) int {
	return fn()
}
// 定义函数类型。
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}
func test2(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s, x)
}
func main() {
	fmt.Println("===========函数定义")
	functionDefinition()
	fmt.Println("===========参数")
	paramDemo()
}

func paramDemo() {
	s := []int{1, 2, 3}
	res := test2("sum: %d", s...)    // slice... 展开slice
	println(res)
}

func functionDefinition() {
	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(s1, s2)
}
