package engine

// 定义结构体

// 发送请求的结构体，url为请求的连接，parser 为返回内容调用的回调函数
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// 请求结果的结构体，返回的结果，url需要再发送请求去获取内容  还有一些内容需要存放进去Item
type ParseResult struct {
	Requests []Request
	Item     []interface{}
}

// 空解析方法，用于测试，某些parser还没完成的时候 后续会不在用
func NilParseFunc(body []byte) ParseResult {
	return ParseResult{}
}
