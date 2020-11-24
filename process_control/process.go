package main

import "fmt"

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
 */
func main() {
	fmt.Println("==========if============")
	ifDemo()
	fmt.Println("===========switch==============")
	switchDemo()
	fmt.Println("==============Type Switch===========")
	typeSwitchDemo()

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
