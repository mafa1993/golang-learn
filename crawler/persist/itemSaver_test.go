package persist

import (
	"crawler/model"
	"testing"
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
	Save(data)
}
