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
	slice()
}

/**
 * 使用指针传递数组，从而实现修改数组值
**/
func changeArr(arr *[3]int) {
	arr[0] = 100
	fmt.Println("函数内", arr)
}

// 切片
func slice() {
	a := [...]int{1, 2, 3, 4, 5}
	s := a[1:3]        // 切片，如果改变s的值，a会跟着变，s为a的一个视图, s为a某一部分的引用
	fmt.Println(s)     // 包含2，不包含4，打印出2,3
	fmt.Println(a[1:]) // 从位置1到最后
	fmt.Println(a[:3]) // 从头到位置3
	fmt.Println(a[:])  // 全部
	s[0] = 1000        // 修改s，a也会变，将s传给函数进行修改，为引用传值，s和a都会变化
	a[2] = 99          // 修改a，s也会变
	fmt.Println("s", s)
	fmt.Println("a", a)

	// 切片可以reslice
	s3 := a[1:4]
	s3 = s3[1:2] // 从上一个s3中取1:2
	fmt.Println(s3)

	// 切面扩展
	arr := [...]int{1, 2, 3, 4, 5}
	s1 := arr[2:4]
	fmt.Println("s1", s1)
	s2 := s1[1:3]
	fmt.Println(s2) // 4，5 对应的为S1的4，s1中没有5，去arr中找，取到4,5  如果超出了arr的长度，会报错，此为切片的扩展
	// len(切片的长度) cap（从切片开头到arr最后一个元素的长度），不超过cap都可以扩展

	// s[i] 不可以超越len[s]  向后扩展不可超过cap(s)

	var ab []int
	fmt.Println(cap(ab))

	s4 := append(s1, 11)
	s5 := append(s2, 13)
	fmt.Println("s4,s5", s4, s5) // s4,s5 [3 4 11] [4 11 13]  s4的改变，导致了s5的改变
	fmt.Println("arr", arr)      // s4的改变导致了 arr的改变
}
