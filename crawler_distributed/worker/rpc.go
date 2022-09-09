package worker

import "crawler/engine"

type CrawlService struct{}

// request 包含一个interface类弄的parse，无法在网络传输，所以要改造
func (c CrawlService) Process(req Request, result *ParseResult) error {
	engineReq := UnSerializeRequest(req)
	
	engineResult, _ := engine.Worker(engineReq)
	// 返回内容
	*result = SerializeResult(engineResult)

	return nil
}
