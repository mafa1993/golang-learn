package engine

// 主任务，队列调度
import (
	"crawler/fetcher"
	"fmt"
	"log"
)

type SimpleEngine struct{}

// 传入多个Request，进行调度
func (simple SimpleEngine) Run(seeds ...Request) { // ...用于接收多个参数，合并成一个slice
	var Requests []Request // 建立一个队列，将所有请求放入

	Requests = append(Requests, seeds...) // 这里... 为展开 传多个参数
	// 和上面一行等效
	// for _, v := range seeds {
	// 	Requests = append(Requests, v)
	// }

	for len(Requests) > 0 {
		Request := Requests[0] // 取出每一个Request进行消费
		Requests = Requests[1:]
		log.Printf("%s", Request.Url)

		rlt, err := simple.Worker(Request)
		if err != nil {
			fmt.Errorf("%s", err)
			continue
		}

		// 将新解析出来的Request对象放入队列
		Requests = append(Requests, rlt.Requests...)

		// item暂时打印
		for _, v := range rlt.Item {
			log.Printf("获取到的itme为%s", v)
		}

	}
}

func (sim SimpleEngine) Worker(request Request) (ParseResult, error) {
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("fetch 出错，msg %s", err)
		return ParseResult{}, err
	}
	rlt := request.ParserFunc(body)

	return rlt, nil
}
