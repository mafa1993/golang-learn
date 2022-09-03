package controller

import (
	"context"
	"crawler/engine"
	"fmt"
	"frontend/model"
	"frontend/view"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/olivere/elastic.v5"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

// 请求的连接为127.0.0.1:8080/?q=xxx&from=xx
//serveHTTP serverHandle必须实现的方法
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 获取参数q
	q := strings.Trim(req.FormValue("q"), " ")

	// 获取参数from
	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		log.Println(err)
		from = 0
	}
	// 向响应输出内容
	fmt.Fprintf(w, "q=%s,from=%d", q, from)
	var page model.SearchResult
	_, page = h.getSearchResult(q, from)

	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func CreateSearchResultHandler(template string) SearchResultHandler {
	// client, err := elastic.NewClient(elastic.SetSniff(false))
	// if err != nil {
	// 	log.Println(err)
	// }
	return SearchResultHandler{
		view: view.CreateSearchResultView(template),
		//client,
		client: &elastic.Client{},
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (error, model.SearchResult) {
	var result model.SearchResult

	q = rewriteQueryString(q)
	resp, err := h.client.Search("dating_profile").Query(elastic.NewQueryStringQuery(q)).From(from).Do(context.Background())

	if err != nil {
		return err, result
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start-len(result.Items)
	result.NextFrom = result.Start+len(result.Items)

	return nil, result

	// for _, v := range resp.Each(reflect.ValueOf(engine.Item)) {
	// 	item := v.(engine.Item)
	// }
}

// 搜索时 将xx：xx 转换为payload.xx:xx
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z]|[a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
