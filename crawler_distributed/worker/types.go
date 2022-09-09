package worker

import (
	"crawler/engine"
	"errors"

	"github.com/thehappymouse/ccmouse/crawler/zhengai/parser"
)

// 定义worker部分的服务等

type SerializedParser struct {
	Name string      // 函数名
	Args interface{} // 函数参数
}

// 服务
type CrawlService struct{}

// 在进行传输的时候需要进行序列化
type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

// 序列化Request，非分布式爬虫的时候，Parse是一个函数，无法进行传输，所以进行序列化
func SerializedRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

// 解析后的结果也是一样，包含函数，需要序列化传输
func SerializeResult(r engine.ParseResult) ParseResult {
	var rlt ParseResult
	rlt.Items = r.Item
	for _, v := range r.Requests {
		rlt.Requests = append(rlt.Requests, SerializedRequest(v))
	}
	return rlt
}

// 反序列化Request
func UnSerializeRequest(r Request) engine.Request {
	return engine.Request{
		Url:    r.Url,
		Parser: UnSerializeParse(r.Parser),
	}
}

// 反序列化parserFunc
// 1 把所有的函数放到一个map里  2 使用switch
func UnSerializeParse(p SerializedParser) engine.Parser {
	switch p.Name {
	case "ParseCity":
		return engine.CreateFuncParserFunc(parser.ParseCity, p.Name)
	case "ParseCityList":
		return engine.CreateFuncParserFunc(parser.ParseCityList, p.Name)
	case "ProfileParser":
		if userName, ok := p.Args.(string); ok {
			return parser.CreateProfileParser(userName), nil
		} else {
			return nil
		}

	case "NilParser":
		return engine.NilParseFunc{}
	default:
		return nil
	}
}

func UnSerializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Item: r.Items,
	}

	for _, v := range r.Requests {
		result.Requests = append(result.Requests, engine.Request{
			Url:    v.Url,
			Parser: UnSerializeParse(v.Parser),
		})
	}
	return result
}


