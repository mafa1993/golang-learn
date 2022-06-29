/*
 * :Date: 2022-06-23 21:12:30
 * :LastEditTime: 2022-06-23 22:04:48
 * :Description:
 */
package main

import (
	"fmt"
	"io"
	"learn/retriever"
	re "learn/retriever"
)

type Retriever interface {
	Get(string) string
}

type Poster interface {
	Post(url string, data map[string]string) string
}

func main() {
	var r Retriever         // Retriever接口
	r = re.Retriever{"abc"} //Retriever结构体实现了Retriever接口 可以当做Retriever接口传入
	fmt.Println(download(r))
	inspect(r)
	r = re.RetrieverT{}
io.ReadWriteCloser
	fmt.Println(11)
	fmt.Printf("%T %v", r, r) // 查看r里面的内容  retriever.RetrieverT { 0s}
	//fmt.Println(download(r))

	/**
	如果实现接口的时候，传的是指针，这里需要写为
	 r = &re.RetrieverT{}

	 func (r *RetrieverT)Get(s string) string {// 函数体}

	 // 打印出来r的类型为*retriever.RetrieverT 值为&{ 0s}
	**/
	inspect(r)

	// type assertion  变量.(类型) 如果类型正确，就可以拿到接口内部的呢绒，如果类型错误会painc, 可以通过接收第二个值
	if realR, ok := r.(retriever.RetrieverT); ok {
		// 使用if 防止panic的出现
		fmt.Println("类型断言", realR.TimeOut)
	}

	fmt.Println(session(re.RetrieverT{}))

}

func inspect(r Retriever) {

	switch v := r.(type) {
	case retriever.Retriever:
		fmt.Println(v.Abc)
	case retriever.RetrieverT:
		fmt.Println(v.TimeOut)
	}
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

// win下执行报错go 1.15.5 amd64  原因暂未找到
// src\bufio\scan.go:184:70: +copy(s.buf, s.buf[s.start:s.end]) evaluated but not used   scan包里多了一个+号。

func post(poster Poster) {
	poster.Post("http://www.baidu.com", map[string]string{"name": "1"})
}

// 建立一个接口，里面有Retriever、Poster 以及一些其他方法, 它既可以调用Retriever的方法，也可以调用Poster的方法
type RetrieverPost interface {
	Retriever
	Poster
	//	Connect(host string)
}

func session(s RetrieverPost) string {
	s.Post("http://www.baidu.com", map[string]string{
		"contents": "abc",
	})
	return s.Get("http://www.baidu.com")
}
