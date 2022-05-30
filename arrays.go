package main

import "fmt"

func main() {
	var arr1 [4]int           // 定义一个长度为4的数组，值类型为int
	arr2 := [3]int{1, 2, 3}   // 定义一个长度为3，值类型为int，并赋为1,2,3
	arr3 := [...]int{2, 3, 4} // 定义一个数组，值类型为int，值为2,3,4，使用编译器去计算数组长度

	var arr4 [2][3]int // 二维数组定义，2行3列
	var arr5 [2][3][4]int

	fmt.Println(arr1, arr2, arr3, arr4, arr5)
	// 数组遍历

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	// 使用range遍历， 只接收一个range的返回值，为键，如果接收两个值，kv都接收，只用k，可以不用_来接收v，如果想只用v，需要使用_来接收k参数
	for i := range arr2 {
		fmt.Println(arr2[i])
	}

	for k, v := range arr2 {
		fmt.Println(k, v)
	}

	fmt.Println("改变前", arr2)
	changeArr(&arr2)
	fmt.Println("函数外arr", arr2)

}

/**
 * 使用指针传递数组，从而实现修改数组值
**/
func changeArr(arr *[3]int) {
	arr[0] = 100
	fmt.Println("函数内", arr)
}
