package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	// 200状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println("响应出错", resp.StatusCode)
		return
	}

	//all, err := ioutil.ReadAll(resp.Body) // read了以后就到了结尾，不能再进行操作
	e := judgeEncoding(resp.Body)
	utf8_reader := transform.NewReader(resp.Body, e.NewDecoder()) // gbk 转成utf-8 NewDecoder() 用什么编码作为解码
	all, err := ioutil.ReadAll(utf8_reader)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
	printCityList(all)
}

/**
判断编码
*/
func judgeEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024) // 从r上创建一个新的reader，并且只取前1024个字符，如果不这样做，一个reader被读取以后，不能再被读取，会导致缺失前1024个字节
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "") // 第一个参数为前1024个字节
	//fmt.Println(e, name, certain)  // 输出&{UTF-8} utf-8 false
	return e
}

/**
 * 获取城市列表
 */
func printCityList(content []byte) {
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/jinan" data-v-f53df81a>济南</a>
	//rex := regexp.MustCompile(`<a[ a-z_A-Z0-9\-]+?href="([a-zA-Z0-9]+)"[^>]+>([^<]+)<`) // [^>]> 这个实现了非贪婪匹配
	rex := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9A-Z]+)"[^>]+>([^<]+)<`) // 应该为470个  494的话，下面有几个推荐城市重复
	matchs := rex.FindAllSubmatch(content, -1)
	for _, v := range matchs {
		fmt.Printf("连接: %s  城市: %s", v[1], v[2])
		fmt.Println()
	}
	fmt.Println(len(matchs))
}
