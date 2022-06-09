package main

import "fmt"

func main() {
	// 创建切片
	var s []int // 在go中，变量声明以后就会有个初始值，切片的初始值为nil

	printSlice(s)

	s1 := []int{2, 3}
	printSlice(s1)

	s2 := make([]int, 16) // 创建长度16的slice
	printSlice(s2)
	s3 := make([]int, 10, 30) // 创建长度为10，cap为30的切片
	printSlice(s3)

	fmt.Println("slice 复制=======")

	copy(s2, s1) // 把s1 拷贝到s2里
	fmt.Println(s2)
	printSlice(s2)

	fmt.Println("删除")
	// 利用append 进行删除
	s2 = append(s2[:1], s2[2:]...) // 删除s2的第二个元素  ... 和php的相同作用，不过php的放在前面，这个放在后面
	fmt.Println(s2)

}

func printSlice(s []int) {
	fmt.Printf("len %d  cap %d \n", len(s), cap(s))
}

 