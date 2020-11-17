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

}

//代码会引发panic
/*
在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间
 */
func testPanic() {
	//panic: runtime error: invalid memory address or nil pointer dereference
	var a *int
	*a = 100
	fmt.Println(*a)

	//panic: assignment to entry in nil map
	var b map[string]int
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
