package main

import (
	"bufio"
	"errors"
	"fmt"
	fic "learn/defererror/Fic"
	"os"
)

func main() {
	tryDefer()
	writeFile2("fic.txt")
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

// 演示panc
func writeFile2(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	/*
				 O_RDONLY：只读模式打开文件；
		O_WRONLY：只写模式打开文件；
		O_RDWR：读写模式打开文件；
		O_APPEND：写操作时将数据附加到文件尾部（追加）；
		O_CREATE：如果不存在将创建一个新文件；
		O_EXCL：和 O_CREATE 配合使用，文件必须不存在，否则返回一个错误；
		O_SYNC：当进行一系列写操作时，每次都要等待上次的 I/O 操作完成再进行；
		O_TRUNC：如果可能，在打开时清空文件
	*/

	if err != nil {
		/*  方法一：输出返回
		fmt.Println(err)
		return
		*/

		// 自定义error，也可以实现Error接口
		error := errors.New("custom error")
		fmt.Println(error)

		// 方案二
		// os.fileopen里面写的，会返回一个*patherror，这里断言判断，如果不是就 退出
		if pathErr, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathErr.Path, pathErr.Err)
			return
		}

		// pathErr的作用域在if里，不能在外部使用

		panic(err) // 会导致程序直接挂掉
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fic.Fic()
	fmt.Fprintln(writer, f())
}
