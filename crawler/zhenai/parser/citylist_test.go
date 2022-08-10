package parser

import (
	"io/ioutil"
	"testing"
)

func TestCityListParser(t *testing.T) {
	// contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	// if err != nil {
	// 	panic(err)
	// }
	// ioutil.WriteFile("./contents.html", contents, 0666)  //将测试数据写入

	content, _ := ioutil.ReadFile("./contents.html")
	res := CityListParser(content)

	resultSize := 470

	if len(res.Item) != resultSize {
		t.Errorf("数量错误,得到的为%d", len(res.Item))
	}

	if len(res.Requests) != resultSize {
		t.Errorf("数量错误,得到的为%d", len(res.Requests))
	}

	// for _, v := range res.Item {
	// 	t.Errorf("%s", v)
	// }

}
