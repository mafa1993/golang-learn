package persist

import (
	"context"
	"crawler/model"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	data := model.Profile{
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
	}
	id, err := Save(data)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	// 从es中获取数据，查看数据是否一致
	resp, err := client.Get().Index("zhenai").Type("doc").Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	var data_de model.Profile
	err = json.Unmarshal(*resp.Source, &data_de)
	if err != nil {
		panic(err)
	}
	t.Logf("%v", data_de)
	if data == data_de {
		t.Errorf("%v,%v", data, data_de)
	}
}
