package view

import (
	"html/template"
	"os"
	"testing"

	"crawler/engine"
	abc "crawler/model"
	"frontend/model"
)

func TestSearch(t *testing.T) {
	template := template.Must(template.ParseFiles("template.html")) // template.Must 对errorjinxing panic

	// 两个参数  一返回的数据存放，二 需要解析的数据
	item := engine.Item{
		Url: "x",
		Id:  "x",
		Payload: abc.Profile{
			Id:         "123456",
			Name:       "test",
			Gender:     "男",                            // 职业
			Age:        "13",                           // 年龄
			Height:     "167",                          // 身高
			Weight:     "65kg",                         // 体重
			Income:     "3000-4000",                    // 收入
			Marriage:   "未婚",                           // 是否已婚
			Education:  "本科",                           // 是否已婚
			Occupation: "it",                           // 职业
			Hokou:      "山东",                           // 籍贯
			Xinzuo:     "天蝎座",                          // 星座
			House:      "无",                            // 房
			Car:        "无",                            // 车
			Address:    "山东济南",                         // 地址
			Photos:     []string{"safs/safd/asdf.png"}, // 照片
			Commit:     "",                             // 备注
		},
	}
	data := model.SearchResult{
		Hits:     1,
		Start:    0,
		Query:    "abc",
		PrevFrom: 0,
		NextFrom: 0,
		//
		//Items:    ,
		Items: []engine.Item{
			item,
		},
	}
	out, _ := os.Create("template.test.html")
	// template.Execute(os.Stdout, data)  // 直接打印到标准输出

	template.Execute(out, data)
}
