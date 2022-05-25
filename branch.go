package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

func main() {
	const filenmae = "abc.txt"
	contents, err := ioutil.ReadFile(filenmae)

	// 写法一
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}

	// 写法二 if语句可以写表达式, if语句可以赋值，这里的赋值的作用域只在if语句只内

	if contents, err = ioutil.ReadFile(filenmae); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	// fmt.Println(
	// 	converToBin(0),
	// 	converToBin(2),
	// 	converToBin(5),
	// )

	readFileAsLine(filename)
}

// if语句
func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

// switch 语句
func eval(a, b int, op string) int {
	var result int

	// switch 语句不用加break，默认有break，如果需要不break，就使用fallthrough 让他继续执行
	// switch 后面可以没有表达式，在case里写条件判断
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("参数错误")
	}

	switch {
	case a > 0:
		fmt.Println("a")
	case a < 0:
		fmt.Println("b")
	default:
		fmt.Println("c")
	}
	return result
}

// for 语句
/**
 * 将10进制转换成2进制
 * 循环/2 取余 直到商变为9 ，然后倒转
 */
func converToBin(a int) string {
	// for循环没有括号
	// for 循环可以没有出事表达式，没有判断表达式，没有递增表达式
	var result string
	if a == 0 {
		return strconv.Itoa(a)
	}
	for ; a > 0; a /= 2 {
		lsb := a % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

var filename string = "abc.txt"

func readFileAsLine(filename string) string {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(1 * time.Second)
	}
	return ""
}
