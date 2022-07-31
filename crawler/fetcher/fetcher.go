package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 发送请求的方法，返回获取的内容和error
func Fetch(url string) ([]byte, error) {
	// 三种实现error的方法  1. errors.New  2. fmt.Errorf() 3. 自己实现error的接口
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	// 200状态码
	if resp.StatusCode != http.StatusOK {
		log.Printf("请求出错，链接为%s, http code为%d", url, resp.StatusCode)
		return nil, fmt.Errorf("请求出错，httpcode 为%d", resp.StatusCode)
	}

	//all, err := ioutil.ReadAll(resp.Body) // read了以后就到了结尾，不能再进行操作
	e := judgeEncoding(resp.Body)
	utf8_reader := transform.NewReader(resp.Body, e.NewDecoder()) // gbk 转成utf-8 NewDecoder() 用什么编码作为解码
	all, err := ioutil.ReadAll(utf8_reader)

	if err != nil {
		log.Printf("请求出错，链接为%s, http code为%d", url, resp.StatusCode)
		return nil, err
	}
	return all, nil
}

/**
判断编码
*/
func judgeEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024) // 从r上创建一个新的reader，并且只取前1024个字符，如果不这样做，一个reader被读取以后，不能再被读取，会导致缺失前1024个字节
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "") // 第一个参数为前1024个字节
	//fmt.Println(e, name, certain)  // 输出&{UTF-8} utf-8 false
	return e
}
