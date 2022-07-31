package parser

import (
	"crawler/engine"
	"fmt"
	"regexp"
)

/**
 * 获取城市列表
 */
const regx string = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9A-Z]+)"[^>]+>([^<]+)<`

func CityListParser(content []byte) engine.ParseResult {
	var rlt engine.ParseResult
	//<a target="_blank" href="http://www.zhenai.com/zhenghun/jinan" data-v-f53df81a>济南</a>
	//rex := regexp.MustCompile(`<a[ a-z_A-Z0-9\-]+?href="([a-zA-Z0-9]+)"[^>]+>([^<]+)<`) // [^>]> 这个实现了非贪婪匹配
	rex := regexp.MustCompile(regx) // 应该为470个  494的话，下面有几个推荐城市重复
	matchs := rex.FindAllSubmatch(content, -1)
	for _, v := range matchs {
		fmt.Printf("连接: %s  城市: %s", v[1], v[2])
		fmt.Println()
		rlt.Requests = append(rlt.Requests, engine.Request{
			Url:        string(v[1]),
			ParserFunc: engine.NilParseFunc, // 暂时使用Nilparser代替
		})
		rlt.Item = append(rlt.Item, v[2]) // 将城市名放到item中
	}

	return rlt
}
