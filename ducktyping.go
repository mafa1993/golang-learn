/*
 * :Date: 2022-06-23 21:12:30
 * :LastEditTime: 2022-06-23 22:01:57
 * :Description:
 */
package main

import (
	"fmt"
	re "learn/retriever"
)

type Retriever interface {
	Get(string) string
}

func main() {
	var r Retriever // Retriever接口
	r = re.Retriever{"abc"}
	fmt.Println(download(r))

	r = re.RetrieverT{}
	fmt.Println(11)
	fmt.Println(download(r))
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

// win下执行报错go 1.15.5 amd64  原因暂未找到
// src\bufio\scan.go:184:70: +copy(s.buf, s.buf[s.start:s.end]) evaluated but not used
