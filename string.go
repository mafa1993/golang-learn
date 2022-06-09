package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string = "我是xx！"

	fmt.Println(len(str))

	for _, v := range []byte(str) { // utf8 编码输出，可变长度，xx为一个字节，中文三个字节
		fmt.Printf("%X  ", v) //%x  16进制输出
	}

	fmt.Println()
	for i, v := range str { // v为rune类型  int32 4个字节
		fmt.Printf("%d,%X ", i, v) // unicode输出  i不是连续的 是字符开始的位置
	}

	fmt.Println()

	for i, v := range []rune(str) { // 转成rune后进行输出，i是连续的，每个v是一个完整的字符
		fmt.Println(i, string(v))
	}

	// 字符长度
	fmt.Println(utf8.RuneCountInString(str))

	//utf8.DecodeRune的作用是通过传入的utf8字节序列转为一个rune即unicode
	bstr := []byte(str)
	for len(bstr) > 0 {
		fmt.Println(bstr)
		a, b := utf8.DecodeRune(bstr) // 一次只会解析一个rune长度的byte
		fmt.Printf("%c,%d ", a, b)    // %c输出一个字符
		bstr = bstr[b:]
	}
}

// 可以使用range进行遍历
// 使用utf8.RuneCountInString获取字符数，使用len获得字节数
// strings 包字符串操作
/**
 常用字符春操作
 Fields、split、Join   strings.Fields() Golang中的函数用于在unicode.IsSpace定义的一个或多个连续空白字符的每个实例周围拆分给定的字符串，返回str的子字符串切片或如果str仅包含空白的空白切片
 Contains、Index
 ToLower,ToUpper
 Trim、TrimRight、TrimLeft
**/
