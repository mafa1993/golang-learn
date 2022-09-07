package main

import (
	"fmt"
	"runtime/debug"
)

// var aa int = 1
// var bb string = "函数外定义变量  包内可见"

// // 函数外不能使用:=
// var (
// 	cc string = "使用var定义多个变量，只写一次var"
// 	dd int    = 1
// )

type a struct {
	e string
}

func (ab a) Error() string {
	debug.PrintStack()
	fmt.Println("调用error方法")
	return ab.e
}

func main() {
	// fmt.Println("hello World")
	// variable()
	// fmt.Println(aa, bb, cc, dd)
	fmt.Println("%s", a{"abc"})

	// euler()

	// transType()
	// consts()

}

// func variable() {
// 	var a int // 先定义名称  再定义类型   例如定义一个变量一般会先想到名字  再想到类型
// 	var s string
// 	fmt.Println(a, s) // int 的初始值是0  string的初始值时"" 在c语言中值不固定，在java中 不管任何类型初始值都为null

// 	var b, c int = 4, 5 // 可以一次定义多个变量，变量必须使用，不使用会报错   避免定义无用变量

// 	var str = "abc" // 可以不定义类型，会自己推断

// 	fmt.Println(b, c, str)

// 	var d, e, f = 1, "e变量值", true // 可以同时给不同类型的变量赋值，不能多个变量定义不同的类型并赋初始值  var a int, b int 这样是不行的

// 	g, h := true, 10 // 初次定义使用:= 推断类型，代替var ,:= 只能在函数内使用，不能在函数外使用

// 	fmt.Println(d, e, f, g, h)

// 	fmt.Printf("%d %q", a, s) // %q和%s区别： %q 会打印出"" 对字符串进行包裹
// }

// /**
//  * 使用go的复数类型 验证欧拉公式
// **/
// func euler() {
// 	//c := 3 + 4i
// 	//fmt.Println("欧拉公式验证",cmplx.Abs(c))
// 	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)

// 	fmt.Println(cmplx.Exp(1i*math.Pi) + 1) // cmplx.Exp表示E的几次方

// }

// /**
//  *  类型转换测试
// **/
// func transType() {
// 	var (
// 		a, b int = 3, 4
// 		c    int
// 	)
// 	//c = math.Sqrt(a*a+b*b); 这里会报错  math.Sqrt返回值为float，复制给int类型会报错，Sqrt的参数要求为float，使用int会报错
// 	c = int(math.Sqrt(float64(a*a + b*b)))
// 	fmt.Println(c)
// }

// /**
//  * 常量定义
//  **/
// func consts() {
// 	const filename string = "abc"
// 	const a, b int = 3, 4
// 	fmt.Println(filename, a, c)
// }
