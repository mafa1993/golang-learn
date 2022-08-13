package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// 增加速率限制
var rateLimite = time.Tick(100 * time.Millisecond)

// 发送请求的方法，返回获取的内容和error
func Fetch(url string) ([]byte, error) {
	<-rateLimite // fetch 争抢rateLimite管道，形成速率限制
	log.Printf("获取的连接%s", url)
	//todo 接口请求有频率限制，暂时无解。
	// 三种实现error的方法  1. errors.New  2. fmt.Errorf() 3. 自己实现error的接口
	//resp, err := http.Get(url)
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:103.0) Gecko/20100101 Firefox/103.0")
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 200状态码
	if resp.StatusCode != http.StatusOK {

		log.Printf("请求出错，链接为%s, http code为%d", url, resp.StatusCode)
		return nil, fmt.Errorf("请求出错，httpcode 为%d", resp.StatusCode)
	}

	//all, err := ioutil.ReadAll(resp.Body) // read了以后就到了结尾，不能再进行操作
	content := bufio.NewReader(resp.Body) // 先复制一份再传递  不然peek会改变原始数据
	// 疑问，这里content不就变成4096个字节，为什么后面从它ReadAll是没问题的

	e := judgeEncoding(content)
	utf8_reader := transform.NewReader(content, e.NewDecoder()) // gbk 转成utf-8 NewDecoder() 用什么编码作为解码
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
func judgeEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024) // 从r上创建一个新的reader，并且只取前1024个字符，如果不这样做，一个reader被读取以后，不能再被读取，会导致缺失前1024个字节

	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "") // 第一个参数为前1024个字节
	//fmt.Println(e, name, certain)  // 输出&{UTF-8} utf-8 false
	return e
}
