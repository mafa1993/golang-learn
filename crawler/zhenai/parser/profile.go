package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"log"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/thehappymouse/ccmouse/crawler/engine"
)

func ParseProfile(content []byte, id string) engine.ParseResult {
	json_data, err := simplejson.NewJson(content)
	if err != nil {
		fmt.Errorf("json解析错误%s", err)
	}

	name, _ := json_data.Get("data").Get("nickname").String()
	gender, _ := json_data.Get("data").Get("genderString").String()
	// basicInfo 数量不一定
	basicInfo, _ := json_data.Get("data").Get("basicInfo").StringArray()
	detailsInfo, _ := json_data.Get("data").Get("detailInfo").StringArray()
	// TODO 多张photo
	// photos, _ := json_data.Get("data").Get("photos").Map()
	photo, _ := json_data.Get("data").Get("avatarURL").String()

	var commit string = ""
	if len(basicInfo) < 9 {
		log.Fatalf("缺少信息%s", basicInfo)
		basicInfo = append(basicInfo, detailsInfo...)
		commit = strings.Join(basicInfo, ",") // 解析不了的放入commit
	}
	if len(detailsInfo) < 10 && commit == "" {
		commit = strings.Join(detailsInfo, ",")
	}

	var photoItem []string
	photoItem = append(photoItem, photo)
	// for _, v := range photos {

	// }

	rtn := model.Profile{
		Id:         id,
		Name:       name,
		Gender:     gender,
		Age:        basicInfo[1],
		Height:     basicInfo[3],
		Weight:     basicInfo[4],
		Marriage:   basicInfo[0],
		Education:  basicInfo[8],
		Occupation: basicInfo[7],
		Hokou:      basicInfo[5],
		Xinzuo:     basicInfo[2],
		House:      detailsInfo[5],
		Car:        detailsInfo[6],
		Address:    detailsInfo[1],
		Photos:     photoItem,
		Commit:     commit,
	}

	result := engine.ParseResult{
		// Item: []interface{}{
		// 	rtn,
		// },
		Item: []engine.Item{
			{
				Url:     fmt.Sprintf("https://album.zhenai.com/u/%s", id),
				Payload: rtn,
				Id:      id,
			},
		},
		Requests: []engine.Request{
			{
				Parser: CreateProfileFunc(id),
			},
		},
	}
	// rlt.Requests = append(rlt.Requests, engine.Request{
	// 	Url: string(v[1]),
	// 	//ParserFunc: engine.NilParseFunc, // 暂时使用Nilparser代替
	// 	//ParserFunc: func(c []byte) engine.ParseResult { return ParseProfile(c, string(v[2])) },
	// 	//ParserFunc: CityParser,
	// 	Parser: engine.CreateFuncParserFunc(CityParser,"CityParser"),
	// })
	return result
}

// 使用正则匹配的，发现有反爬
// func ParseProfile(content []byte) engine.ParseResult {
// 	ioutil.WriteFile("./a.json", content, 0666)
// 	var profile_pre *regexp.Regexp = regexp.MustCompile(`"m-btn purple">(.+)[^<]`)
// 	var profile_pre2 *regexp.Regexp = regexp.MustCompile(`"m-btn pink">(.+)[^<]`)
// 	var id_pre *regexp.Regexp = regexp.MustCompile(`ID：([0-9]*)`)
// 	var name_pre *regexp.Regexp = regexp.MustCompile(`nickName.*>(.*)[^<]`)
// 	var gender_pre *regexp.Regexp = regexp.MustCompile(`([男女])士征婚`)
// 	var photo_pre *regexp.Regexp = regexp.MustCompile(`<div.*href="([^"])"class="photoItem`)
// 	rlt := profile_pre.FindAllSubmatch(content, -1)
// 	rlt2 := profile_pre2.FindAllSubmatch(content, -1)
// 	photo := photo_pre.FindAllSubmatch(content, 3)
// 	gender := gender_pre.FindSubmatch(content)
// 	for _, v := range rlt {
// 		fmt.Println(v)
// 	}
// 	if len(rlt) < 9 {
// 		panic("错误")
// 	}
// 	ID := id_pre.FindSubmatch(content)
// 	if ID == nil {
// 		panic("未找到id")
// 	}
// 	Name := name_pre.FindSubmatch(content)
// 	if Name == nil {
// 		panic("name 没找到")
// 	}
// 	var photos []string
// 	for _, v := range photo {
// 		photos = append(photos, string(v[1]))
// 	}

// 	rtn := model.Profile{
// 		Id:         string(ID[1]),
// 		Name:       string(Name[1]),
// 		Gender:     string(gender[1]),
// 		Age:        string(rlt[1][2]),
// 		Height:     string(rlt[1][4]),
// 		Weight:     string(rlt[1][5]),
// 		Marriage:   string(rlt[1][1]),
// 		Education:  string(rlt[1][9]),
// 		Occupation: string(rlt[1][8]),
// 		Hokou:      string(rlt[1][6]),
// 		Xinzuo:     string(rlt[1][3]),
// 		House:      string(rlt2[1][6]),
// 		Car:        string(rlt2[1][7]),
// 		Address:    string(rlt[1][5]) + string(rlt2[1][2]),
// 		Photos:     photos,
// 	}

// 	result := engine.ParseResult{
// 		Item: []interface{}{
// 			rtn,
// 		},
// 	}
// 	return result
// }

// profile不通用，多一个参数，需要返回，单独建立
type ProfileParser struct {
	username string
}

// funcparser 实现interface
func (f *ProfileParser) Parse(contents []byte) engine.ParseResult {
	return ParseProfile(contents, f.username)
}

func (f *ProfileParser) Serialize() (name string, args interface{}) {
	return "ParseProfile", f.username
}

func CreateProfileFunc(username string) *ParseProfile {
	return &ParseProfile{
		name: username,
	}
}
