package main

import (
	"fmt"
	"time"
)

/**
条件语句if:
if 语句 由一个布尔表达式后紧跟一个或多个语句组成
格式：
	if 布尔表达式 {
		//在布尔表达式为 true 时执行
	}
    • 可省略条件表达式括号。
    • 持初始化语句，可定义代码块局部变量。
    • 代码块左 括号必须在条件表达式尾部。
PS:不支持三元操作符(三目运算符) "a > b ? a : b"。

条件语句switch:
switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上直下逐一测试，直到匹配为止。
Golang switch 分支表达式可以是任意类型，不限于常量。可省略 break，默认自动终止.
语法如下:
	switch var1 {
		case val1:
			...
		case val2:
			...
		default:
			...
	}
PS:
变量 var1 可以是任何类型，而 val1 和 val2 则可以是同类型的任意值。
类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。
Type Switch:
被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型
语法格式如下：
	switch x.(type){
		case type:
		   statement(s)
		case type:
		   statement(s)
		//可以定义任意个数的case
		default: //可选
			statement(s)
	}

条件语句select
select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行
select 语句的语法如下:
	select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
    // 你可以定义任意数量的 case
	default : //可选
	statement(s);
	}
PS:
	每个case都必须是一个通信
    所有channel表达式都会被求值
    所有被发送的表达式都会被求值
    如果任意某个通信可以进行，它就执行；其他被忽略。
    如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
    否则：
    如果有default子句，则执行该语句。
    如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。
select由比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作
在一个select语句中，Go会按顺序从头到尾评估每一个发送和接收的语句。
如果其中的任意一个语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。
如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有两种可能的情况：
①如果给出了default语句，那么就会执行default的流程，同时程序的执行会从select语句后的语句中恢复。
②如果没有default语句，那么select语句将被阻塞，直到至少有一个case可以进行下去。
 */
func main() {
	fmt.Println("==========if============")
	ifDemo()
	fmt.Println("===========switch==============")
	switchDemo()
	fmt.Println("==============Type Switch===========")
	typeSwitchDemo()
	fmt.Println("==========select==============")
	selectDemo()
	fmt.Println("============select time out ==========")
	selectTimeOutDemo()

}
var resChan = make(chan int)
func selectTimeOutDemo() {
	select {
	case data := <-resChan:
		doData(data)
	case <-time.After(time.Second * 3):
		fmt.Println("request time out")
	}
}
func doData(data int) {
	//...
	fmt.Println("---doData---")
}

func selectDemo() {
	var c1, c2, c3 chan int //一个可以发送 int 类型数据的 channel 一般写为 chan int
	var i1, i2 int
	select { //不停的在这里检测
	case i1 = <-c1: //检测有没有数据可以读, 如果cl成功读取到数据，则进行该case处理语句 //阻塞接收数据
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2: //检测有没有可以写，如果成功向c2写入数据，则进行该case处理语句
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3 //非阻塞接收数据
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default: //假如没有default，那么在以上两个条件都不成立的情况下，就会在此阻塞//一般default会不写在里面，select中的default子句总是可运行的，因为会很消耗CPU资源
		//如果以上都没有符合条件，那么则进行default处理流程
		fmt.Printf("no communication\n")
	}
}

func typeSwitchDemo() {
	var x interface{}
	//写法一：
	switch i := x.(type) { // 带初始化语句
	case nil:
		fmt.Printf(" x 的类型 :%T\r\n", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
	//写法二
	var j = 0
	switch j {
	case 0:
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("j=====>def")
	}
	//写法三
	var k = 0
	switch k {
	case 0:
		fmt.Println("fallthrough")
		fallthrough
		/*
		   Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		   而如果switch没有表达式，它会匹配true。
		   Go里面switch默认相当于每个case最后带有break，
		   匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		   但是可以使用 fallthrough 强制执行后面的case代码。
		*/
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("k=====>def")
	}
	//写法三
	var m = 0
	switch m {
	case 0, 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("m=====>def")
	}
	//写法四
	var n = 0
	switch { //省略条件表达式，可当 if...else if...else
	case n > 0 && n < 10:
		fmt.Println("i > 0 and i < 10")
	case n > 10 && n < 20:
		fmt.Println("i > 10 and i < 20")
	default:
		fmt.Println("n=====>def")
	}
}

func switchDemo() {
	var grade = "B"
	var marks = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 60, 70: grade = "C"
		default: grade = "D"
	}
	//如果switch没有表达式，它会匹配true
	switch {
	case grade == "A":
		fmt.Println("优秀!")
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" )
	}
	fmt.Printf("你的等级是 %s\n", grade )
}

func ifDemo() {
	x := 0
	if n := "abc";x > 0 {
		println(n[2])
	} else if x < 0 {
		println(n[1])
	} else{
		println(n[0])
	}
}
