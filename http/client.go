package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	text, err := httputil.DumpResponse(res, true) // 第二个参数为是否dump body,text为byte类型

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", text)

	client()
}

func client() {
	request, err := http.NewRequest(http.MethodGet, "https://www.gitee.com", nil) // 第三个参数为request.body

	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 10; EVR-AL00 Build/HUAWEIEVR-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.186 Mobile Safari/537.36 baiduboxapp/11.0.5.12 (Baidu; P1 10)") // 增加请求头，user-agent

	// client里可以检测是否有重定向，放cookie等
	customClient := http.Client{
		// via会保存所有的重定向路径，如果返回nil 就继续redirect  如果返回错误 终止重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("重定向", req)
			return nil
		},
	}
	//res, err := http.DefaultClient.Do(request) // 使用默认的客户端发出请求
	res, err := customClient.Do(request) // 使用自定义client 发出请求
	defer res.Body.Close()               // 关闭请求
	//testres(res)

}

// 对response 进行调试
func testres(res *http.Response) {
	rlt, err := httputil.DumpResponse(res, true)
	if nil != err {
		panic(err)
	}

	fmt.Printf("%s\n", rlt)

}
