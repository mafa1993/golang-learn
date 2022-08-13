package parser

// 用于定义各个parser ,包含入口页  城市页  个人信息页

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
	//var i int = 0
	for _, v := range matchs {
		//name := string(v[2])  //v其实为引用，直接使用v[2],在parserfunc 执行的时候，name为同一个

		fmt.Println()
		rlt.Requests = append(rlt.Requests, engine.Request{
			Url: string(v[1]),
			//ParserFunc: engine.NilParseFunc, // 暂时使用Nilparser代替
			//ParserFunc: func(c []byte) engine.ParseResult { return ParseProfile(c, string(v[2])) },
			ParserFunc: CityParser,
		})
		rlt.Item = append(rlt.Item, v[2]) // 将城市名放到item中
		// i++
		// if i > 0 {
		// 	break
		// }
	}

	return rlt
}
