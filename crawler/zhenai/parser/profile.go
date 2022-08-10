package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"io/ioutil"
	"regexp"
)

func ParseProfile(content []byte) engine.ParseResult {
	ioutil.WriteFile("./a.json", content, 0666)
	var profile_pre *regexp.Regexp = regexp.MustCompile(`"m-btn purple">(.+)[^<]`)
	var profile_pre2 *regexp.Regexp = regexp.MustCompile(`"m-btn pink">(.+)[^<]`)
	var id_pre *regexp.Regexp = regexp.MustCompile(`ID：([0-9]*)`)
	var name_pre *regexp.Regexp = regexp.MustCompile(`nickName.*>(.*)[^<]`)
	var gender_pre *regexp.Regexp = regexp.MustCompile(`([男女])士征婚`)
	var photo_pre *regexp.Regexp = regexp.MustCompile(`<div.*href="([^"])"class="photoItem`)
	rlt := profile_pre.FindAllSubmatch(content, -1)
	rlt2 := profile_pre2.FindAllSubmatch(content, -1)
	photo := photo_pre.FindAllSubmatch(content, 3)
	gender := gender_pre.FindSubmatch(content)
	for _, v := range rlt {
		fmt.Println(v)
	}
	if len(rlt) < 9 {
		panic("错误")
	}
	ID := id_pre.FindSubmatch(content)
	if ID == nil {
		panic("未找到id")
	}
	Name := name_pre.FindSubmatch(content)
	if Name == nil {
		panic("name 没找到")
	}
	var photos []string
	for _, v := range photo {
		photos = append(photos, string(v[1]))
	}

	rtn := model.Profile{
		Id:         string(ID[1]),
		Name:       string(Name[1]),
		Gender:     string(gender[1]),
		Age:        string(rlt[1][2]),
		Height:     string(rlt[1][4]),
		Weight:     string(rlt[1][5]),
		Marriage:   string(rlt[1][1]),
		Education:  string(rlt[1][9]),
		Occupation: string(rlt[1][8]),
		Hokou:      string(rlt[1][6]),
		Xinzuo:     string(rlt[1][3]),
		House:      string(rlt2[1][6]),
		Car:        string(rlt2[1][7]),
		Address:    string(rlt[1][5]) + string(rlt2[1][2]),
		Photos:     photos,
	}

	result := engine.ParseResult{
		Item: []interface{}{
			rtn,
		},
	}
	return result
}
