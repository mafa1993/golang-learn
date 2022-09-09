package engine

import (
	"crawler/fetcher"
	"log"
)

func (sim SimpleEngine) Worker(request Request) (ParseResult, error) {
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("fetch 出错，msg %s", err)
		return ParseResult{}, err
	}
	rlt := request.ParserFunc(body)
	return rlt, nil
}
