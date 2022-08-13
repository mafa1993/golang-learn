package parser

import (
	"crawler/engine"
	"regexp"
)

var cityRe *regexp.Regexp = regexp.MustCompile(`"http://album.zhenai.com/u/([0-9]+)"`)

func CityParser(content []byte) engine.ParseResult {
	rlt := cityRe.FindAllSubmatch(content, -1)
	var rtn engine.ParseResult
	for _, v := range rlt {
		rtn.Item = append(rtn.Item, v[1])
		rtn.Requests = append(rtn.Requests, engine.Request{
			Url: `https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=` + string(v[1]) + `&_=1660138857621&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2Fbd5e658a-a3da-4380-b0ea-7249493b2ad0%2F0%2F0%2F1757971071&data=eyJ2IjoiSEJOTWgrMTM3Ym9qMkJQL0NaZS93Zz09Iiwib3MiOiJ3ZWIiLCJpdCI6MjAyOCwidCI6ImpTU0JDVUl3SlY5S2pObndlemg1eHhMVjNUeUFUckxEWXRMZndZS2tZMWJQZm1abGl0VWFkeWVnak5DUlN2UnZsanVzUkRrTDMxNlgxUjBwVktPTk53PT0ifQ%3D%3D&MmEwMD=5ltyrP4ajWHJfXW9kxhr6hVeYECnTtP2PQBLpveeQ2n53URLTOqP21DHSj_A6picEYDoTVXorTo._pkfpaCYZiKgWpaVetLZyiunDQy9IWsqhVBaEuADisYOhiDFfWE1qdwQ00YkDmtLlTR23nQf11cNV3y5qT0ZsueRP0qG4j_9nWPUpdV9N4xyPSNf4zpx65QdXumulgLLUAX4cV1PL1mJumX2cLijKN83WIMD7Txaw30QH0.FQYEshoMPLL6hdaUE5OV.ORWP7qK1bnc4pmdKtuxz_X5FsbIwvZ6t2Dl5Vq4izsrg325DJ7P5zg4y7`,
			//Url:string(v[1]),
			ParserFunc: func(content []byte) engine.ParseResult { // content会在使用的时候传进去，参数二线传进去
				return ParseProfile(content, string(v[1]))
			},
		})
	}
	return rtn
}
