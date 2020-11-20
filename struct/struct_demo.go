package main

import "fmt"

//Go语言中没有“类”的概念，也不支持“类”的继承等面向对象的概念。Go语言中通过结构体的内嵌再配合接口比面向对象具有更高的扩展性和灵活性。
/*
自定义类型：
Go语言中可以使用type关键字来定义自定义类型。我们可以基于内置的基本类型定义，也可以通过struct定义
eg: type MyInt int ////将MyInt定义为int类型,MyInt就是一种新的类型，它具有int的特性
类型别名:
类型别名是Go1.9版本添加的新功能
类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
eg: type TypeAlias = Type  type byte = uint8  type rune = int32 //rune和byte就是类型别名
类型别名只会在编码中存在，编译完成时并不会有
 */
func main() {
	fmt.Println("==============自定义类型和类型别名==============")
	typeAliasTest()
}

//演示自定义类型和类型别名的区别
func typeAliasTest() {
	type newInt int
	type myInt = int

	var a newInt
	var b myInt
	fmt.Printf("type of a:%T\n", a) // type of a:main.newInt
	fmt.Printf("type of b:%T\n", b)  // type of b:int
}
