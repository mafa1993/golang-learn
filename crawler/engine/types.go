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
	Item     []Item
}

// 为了扩展Item，创建一个结构体，Payload里面存返回的内容，id存人的id，防止重复录入，url存这个人的Url
type Item struct {
	Id string 
	Url string 
	Payload interface{}  // 这里payload为interface类型，在用josn.Unmarshal解码时，不一定会转换成什么类型，我们想要将他转换成model.profile
}


// 空解析方法，用于测试，某些parser还没完成的时候 后续会不在用
func NilParseFunc(body []byte) ParseResult {
	return ParseResult{}
}
