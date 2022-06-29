/*
 * :Date: 2022-06-23 21:36:35
 * :LastEditTime: 2022-06-23 21:57:11
 * :Description:
 */
package retriever

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

type RetrieverT struct {
	UserAgent string
	TimeOut   time.Duration // time.Duration 时间 毫秒
	Contents  string
}

// 只要其他类型实现了接口的方法，即实现了接口
func (r RetrieverT) Get(str string) string {
	fmt.Println(str)
	res, err := http.Get(str)
	fmt.Println(1)
	if err != nil {
		panic(err)
	}
	rlt, err := httputil.DumpResponse(res, true) // 输出http的返回 第二个参数为是否输出响应体
	//defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(rlt)
}

func (r RetrieverT) Post(url string, data map[string]string) string {
	r.Contents = data["contents"]
	fmt.Println(r.Contents)
	return "ok"
}
