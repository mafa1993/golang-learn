// 讲解函数式编程

package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ===== 闭包讲解

// 定义函数add，返回值是一个函数,返回的函数的参数为int，返回值为int
func adder() func(int) int {
	var sum int = 0
	return func(i int) int {
		sum += i // sum 作用域链，找外面的sum
		return sum
	}
}

// 闭包end ==========
func main() {
	a := adder() // adder调用一次，固化sum的初始值

	for i := 0; i <= 9; i++ {
		fmt.Println(a(i)) // a为adder返回的函数，调用n次
	}

	a2 := adder2(0)
	var s int
	for i := 0; i <= 9; i++ {
		s, a2 = a2(i) //a2每次都为新的函数返回
		fmt.Println(s, a2)
	}

	fmt.Println("=======斐波那契")
	f := fibnacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("利用Reader ")
	f2 := ff() // 调用ff以后返回了一个Reader  传给printFile，去循环遍历   ff相当于一个生成器
	printFile(f2)

}

// 使用函数式编程实现，使用一个函数来作为sum的记录=======
type baseAdd func(int) (int, baseAdd) // 定义一个baseAdd类型，其本身为函数，参数为int，返回值的第一个int为和，第二个为累加函数
// 利用递归的思想

func adder2(base int) baseAdd {
	return func(v int) (int, baseAdd) {
		return base + v, adder2(base + v)
	}
}

// end =======

// ==== 斐波那契数列(第三次调用是前两次的和) ====
func fibnacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

// ====

// === 使用函数实现接口

// 给斐波那契实现Reader的Read接口，然后使用Scan调用

type intGen func() int // 顶一个类型，本身为函数类型，作为斐波那契的返回值类型

func ff() intGen {
	var a, b int = 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

// 实现read接口, 读一个byte类型，返回读取的长度和错误
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 100 {
		return 0, io.EOF // 大于100 就停止
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p) // 把字符串s转换成reader，放到p byte中
}

func printFile(reader io.Reader) {
	scaner := bufio.NewScanner(reader)
	for scaner.Scan() {
		fmt.Println(scaner.Text())
	}
}
