package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 内置库的使用

type Retriever struct {
	Contents string
}

// 相当于toString
func (r Retriever) String() string {
	return fmt.Sprintf("sss  %s", r.Contents)
}

func main() {
	//如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。
	r := Retriever{Contents: "sdf"}
	fmt.Println(r)

	// fmt.Fprintf() 向io.Write中写, 可以是网络，也可以是文件
	// fmt.Fscanf() 向io.Reader里读，可以是网络io，文件io等

	var str string = `sadfas
	sfs
	
	sf` // ``里面放多行文本
	var filename string = "abc.txt"
	readFileAsLine(filename)
	printFileContent(strings.NewReader(str))  // strings.NewReader()  将字符串转换成reader

}

func readFileAsLine(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	printFileContent(file)

}

func printFileContent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(1 * time.Second)
	}
}
