package main

import (
	_ "learn-go/hello"  // _ 表示只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数
)

func main()  {
	//hello.Print()
}
