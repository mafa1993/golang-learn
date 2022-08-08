package model

type Profile struct {
	Id         string
	Name       string
	Gender     string   // 职业
	Age        int      // 年龄
	Height     int      // 身高
	Weight     int      // 体重
	Income     string   // 收入
	Marriage   string   // 是否已婚
	Education  string   // 是否已婚
	Occupation string   // 职业
	Hokou      string   // 籍贯
	Xinzuo     string   // 星座
	House      string   // 房
	Car        string   // 车
	Address    string   // 地址
	Photos     []string // 照片
}
