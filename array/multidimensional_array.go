package main

import "fmt"

//多维数组
/*
全局
    var arr0 [5][3]int
    var arr1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
    局部：
    a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
    b := [...][2]int{{1, 1}, {2, 2}, {3, 3}} // 第 2 纬度不能用 "..."。 use of [...] array outside of array literal
*/

var arr5 [5][3]int
var arr6 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}

func main() {
	a := [2][3]int{{1, 2},{4, 5, 6}}
	b := [...][2]int{{1, 2},{2, 2},{3, 4}}
	//c := [2][...]int{{1, 2, 3},{3, 4, 5}}  //use of [...] array outside of array literal
	fmt.Println(arr5, arr6)
	fmt.Println(a, b)
	fmt.Println(len(arr6), len(arr6[1]))


	//遍历二维数组
	for k1,v1 := range b {
		for k2,v2 := range v1 {
			fmt.Printf("(%d,%d)=%d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

