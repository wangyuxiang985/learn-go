package main

import "fmt"

//Go语言中的指针不能进行偏移和运算，是安全指针
// &（取地址）和 *（根据地址取值）
//Go语言中的值类型（int、float、bool、string、array、struct）都有对应的指针类型，如：*int、*int64、*string等
/**
ptr := &v    // v的类型为T
v:代表被取地址的变量，类型为T
ptr:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
 */

//当一个指针被定义后没有分配到任何变量时，它的值为 nil

//要分配内存，就new和make。 Go语言中new和make是内建的两个函数，主要用来分配内存
/**
new是一个内置的函数，它的函数签名如下：
func new(Type) *Type
1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值
 */
/**
make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型。
make函数的函数签名如下:
func make(t Type, size ...IntegerType) Type
PS: make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作
 */
/**
new 与 make 区别与联系：
1.二者都是用来做内存分配的。
2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
 */
func main()  {
	a := 10
	b := &a //取变量a的地址，将指针保存到b中
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc0000a0090
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc0000a0090 type:*int
	fmt.Println(&b)                    // 0xc0000ca018
	fmt.Println("===========")
	c := *b //指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
	fmt.Println("------modify")

	d := 10
	modify1(d)
	fmt.Println(d)
	modify2(&d)
	fmt.Println(d)
	fmt.Println("******************")
	nil_demo()
	fmt.Println("&&&&&&&&&&&&&&&&&&&&&&")
	testPanic()
	fmt.Println("--------------------new-----------")
	test_new()
}

func test_new() {
	a := new(int)
	b := new(bool)
	fmt.Printf("%T\n", a) // *int
	fmt.Printf("%T\n", b) // *bool
	fmt.Println(*a)       // 0
	fmt.Println(*b)       // false
}


//代码会引发panic
/*
在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间
 */
func testPanic() {
	//panic: runtime error: invalid memory address or nil pointer dereference
	var a *int  //声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值
	//对a进行初始化操作
	a = new(int)
	*a = 100
	fmt.Println(*a)

	//panic: assignment to entry in nil map
	var b map[string]int //只是声明变量b是一个map类型的变量，需要使用make函数进行初始化操作之后，才能对其进行键值对赋值
	b = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}
//nil值判断
func nil_demo()  {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是%v\n", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

func modify1(x int) {
	x = 100
}
func modify2(x *int)  {
	*x = 100
}
