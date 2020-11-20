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
/*
结构体：
Go语言中通过struct来实现面向对象
结构体定义：使用type和struct关键字来定义结构体
type 类型名 struct {
        字段名 字段类型
        字段名 字段类型
        …
}
1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
3.字段类型：表示结构体字段的具体类型。
ps:只有当结构体实例化时，才会真正地分配内存,必须实例化后才能使用结构体的字段
结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型:
var 结构体实例 结构体类型

创建指针类型的结构体
通过使用new关键字对结构体进行实例化，得到的是结构体的地址,Go语言中支持对结构体指针直接使用.来访问结构体的成员

取结构体的地址实例化
使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作

使用键值对初始化
使用键值对对结构体进行初始化时，键对应结构体的字段，值对应该字段的初始值

使用值的列表初始化
初始化结构体的时候可以简写，也就是初始化的时候不写键，直接写值,需要注意：
    1.必须初始化结构体的所有字段。
    2.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
    3.该方式不能和键值初始化方式混用。

构造函数：
Go语言的结构体没有构造函数，但是可以自己实现；
因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
 */
func newPerson(name, nickName string, age int) *person {
	return &person{
		name:name,
		nickName:nickName,
		age:age,
	}
}
type person struct {
	name,nickName string
	age int
}
func main() {
	fmt.Println("==============自定义类型和类型别名==============")
	typeAliasTest()
	fmt.Println("=============结构体===================")
	personStruct()
	fmt.Println("================匿名结构体=================")
	anonymityStruct()
	fmt.Println("===============指针类型结构体=============")
	pointerStruct()
	fmt.Println("===============取结构体的地址实例化===================")
	adressStruct()
	fmt.Println("============键值对初始化==============")
	keyValueInit()
	fmt.Println("===========使用值的列表初始化==============")
	valueInit()
	fmt.Println("================面试题=================")
	interviewDemo()
	fmt.Println("==========构造函数============")
	constructorDemo()

}

//模拟构造函数
func constructorDemo() {
	p5 := newPerson("哪吒","三太子", 20)
	fmt.Printf("p5=%#v\n",p5)
}

func interviewDemo() {
	m := make(map[string]*person)
	stus := []person{
		{name: "悟空", age: 18},
		{name: "八戒", age: 23},
		{name: "沙僧", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		fmt.Printf("v=%T\n",v)
	}
}

//初始化值的方式初始化结构体
func valueInit() {
	p4 := &person{
		"唐僧",
		"御弟哥哥",
		18,
	}
	p5 := person{
		"唐僧",
		"御弟哥哥",
		18,
	}
	fmt.Printf("p4=%#v\n", p4)
	fmt.Printf("p5=%#v\n", p5)
}

//键值对的方式初始化
func keyValueInit() {
	p3 := person{
		name:"沙僧",
		age: 10,
	}
	fmt.Printf("p3=%#v\n", p3)
}

//取结构体的地址实例化
func adressStruct() {
	p2 := &person{}
	fmt.Printf("p2=%T\n",p2)
	fmt.Printf("p2=%#v\n", p2)
	fmt.Println(p2)
	// p2.name = "猪八戒"其实在底层是(*p2).name = "猪八戒"，这是Go语言帮我们实现的语法糖
	p2.name = "猪八戒"
	p2.age = 0
	fmt.Printf("p2=%#v\n",p2)
}

//指针类型结构体
func pointerStruct() {
	var p1 = new(person)
	//%T:相应值的类型; %#v:相应值
	fmt.Printf("%T\n", p1) //*main.person
	fmt.Printf("p2=%#v\n", p1)  //p2=&main.person{name:"", nickName:"", age:0}
	fmt.Println(p1) //&{  0}
	p1.name = "孙悟空"
	p1.age = 1
	fmt.Println(p1) //&{孙悟空  1}
	fmt.Printf("p1=%#v\n", p1) //p1=&main.person{name:"孙悟空", nickName:"", age:1}
}

//匿名结构体
func anonymityStruct() {
	var student struct{name string; age int}
	student.name = "李四"
	student.age = 20
	fmt.Printf("%#v\n", student)
}

//结构体的简单使用
func personStruct() {
	var xiaowang person
	xiaowang.name = "王五"
	xiaowang.nickName = "老王"
	xiaowang.age = 18
	fmt.Printf("xiaowang=%v\n", xiaowang)
	fmt.Printf("xiaowang=%#v\n", xiaowang)
	fmt.Println(xiaowang.age)
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
