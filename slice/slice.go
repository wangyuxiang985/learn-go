package main

import "fmt"

//slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。
//1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
//2. 切片的长度可以改变，因此，切片是一个可变的数组。
//3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
//4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
//5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
//6. 如果 slice == nil，那么 len、cap 结果都等于 0。

func main()  {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("s1为空")
	}else {
		fmt.Println("s1不为空")
	}
	//2. :=
	s2 := []int{}
	fmt.Println("s2:", s2)
	//3. make()
	var s3 []int = make([]int, 0)
	fmt.Println("s3:", s3)
	//4. 初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println("s4:", s4)
	s5 := []int{1, 2, 3}
	fmt.Println("s5:", s5)

	//5. 从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	//前包后不包
	s6 = arr[1:4]
	fmt.Println("s6:", s6)

	data := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s7 := data[2:4]
	s7[0] += 100
	s7[1] += 200
	fmt.Println("s7:", s7)
	fmt.Println("data:", data)

	s8 := []int{1, 2, 3, 4, 5, 6: 100}
	fmt.Println("s8:", s8, len(s8), cap(s8))

	s9 := make([]int, 3, 5)// 使用 make 创建，指定 len 和 cap 值,若省略cap值则cap=len
	fmt.Println("s9:", s9, len(s9), cap(s9))

	// & 是取地址符号 , 即取得某个变量的地址 , 如 &s10[2] 表示取s10[2]值的内存地址
	// *是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值
	s10 := []int{1, 2, 3}
	fmt.Println("s10:",s10)
	p := &s10[2]
	*p += 100
	fmt.Println("指针操作后s10：", s10)

	// append后得到一个新的对象
	s11 := make([]int, 2)
	fmt.Println("s11:", s11)
	s12 := append(s11, 100)
	fmt.Println("s12:", s12)
	fmt.Printf("%p\n",&s11)
	fmt.Printf("%p\n",&s12)

	//如果append追加超出原 slice.cap 限制，就会重新分配底层数组，一般扩容为原来的2倍
	data1 := []int{0, 1, 2, 3, 4, 10: 0}
	fmt.Println("data1:", data1)
	sliceData1 := data1[:2:3]
	fmt.Println("sliceData1:", sliceData1, len(sliceData1), cap(sliceData1))
	fmt.Println( &data1[0],&sliceData1[0])
	sliceData1 = append(sliceData1, 10, 30)
	fmt.Println("sliceData1:", sliceData1, len(sliceData1), cap(sliceData1))
	fmt.Println( &data1[0],&sliceData1[0])
	//copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠
	//常规slice , data[2:3]，从第2位到第3位（返回1，2），长度len为2， 最大可扩充长度cap为4（6-9）
	//另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8

	//a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x
	data2 := []int{0, 1, 2, 3, 4, 5, 6}
	sData2 := data2[1:3:4]
	fmt.Println("data2:", data2, len(data2), cap(data2))
	fmt.Println("sData2:", sData2, len(sData2), cap(sData2))
}
