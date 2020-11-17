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

}
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
