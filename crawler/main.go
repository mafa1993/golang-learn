package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
)

func main() {
	var seeds []engine.Request

	seeds = []engine.Request{
		{
			Url:        "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.CityListParser,
		},
	}

	engine.Run(seeds...)
}
