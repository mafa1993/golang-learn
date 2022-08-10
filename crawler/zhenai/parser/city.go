package parser

import (
	"crawler/engine"
	"regexp"
)

var cityRe *regexp.Regexp = regexp.MustCompile(`"(http://album.zhenai.com/u/[0-9]+)"`)

func CityParser(content []byte) engine.ParseResult {
	rlt := cityRe.FindAllSubmatch(content, -1)
	var rtn engine.ParseResult
	for _, v := range rlt {
		rtn.Item = append(rtn.Item, v[1])
		rtn.Requests = append(rtn.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: ParseProfile,
		})
	}
	return rtn
}
