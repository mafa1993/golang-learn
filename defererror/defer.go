package main

import (
	"bufio"
	"fmt"
	fic "learn/defererror/Fic"
	"os"
)

func main() {
	tryDefer()

	writeFile("fic.txt")
}

// 打印顺序为3 1 2 先derfer的后执行
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	// 理解在defer在遇到时进行计算
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename) // os.Create创建文件 如果文件已存在，会将文件清空
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file) //创建一个缓冲区，先将内容写到内存
	defer writer.Flush()            // 将缓冲区内容 写入到文件

	f := fic.Fic()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f()) // 将f的返回内容，写入到writer中
	}

}
