package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	q, r := dd(11, 2)
	x, _ := dd(1, 2) // 使用_ 屏蔽不用的返回值，如果接收了不使用  会有编译错误

	fmt.Println(q, r, x)

	err := errorRtn()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(apply(func(a, b int) int {
		return 1
	}, 3, 4)) // 输出：函数名为main.main.func1, args (3,4)1   第一个main为包名  第二个main为函数名，func1 为匿名函数取得名

	fmt.Println(sum(1, 2, 3, 4, 6))
	a,b := 1,2
	swap(&a,&b)
	fmt.Println(a,b)
}

// q,r 会被预先定义好，在函数内不需要再进行声明,q r的作用域在函数内部, 并
func dd(a, b int) (q int, r int) {
	q = a / b // int / int嘚int  不会得到float
	r = a % b
	// 直接return, 会把函数内部的q r 返回
	return
}

func errorRtn() error {
	return fmt.Errorf("abc%s", "出错了")
}

// 参数为函数，以及反射初探
// 参数一 op是个函数，这个函数的参数为int，返回值为int
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()     // 获取op的反射，然后获取反射的指针
	funName := runtime.FuncForPC(p).Name() //通过reflect的ValueOf().Pointer作为入参，获取函数地址、文件行、函数名等信息

	fmt.Printf("函数名为%s, args (%d,%d)", funName, a, b)
	return op(a, b)
}

// 不定传参
func sum(num ...int) int {
	var s int
	for i := range num {
		s += num[i]
	}

	return s
}

/*
 * 函数的用法特点
 1. 返回值类型写在最后
 2. 可返回多个值
 3. 函数作为参数
 4. 没有默认参数，可选参数等 有可变参数列表
*/


// 函数值传递

// 用于两个变量值互换
func swap(a,b *int) {  // 传递a,b的指针
	*a,*b = *b,*a  // a指向的内容改为b指向的内容。。。
}