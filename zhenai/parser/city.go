package parser

import (
	"crawler/engine"
	"regexp"
)

var (
	profileRe  = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/shanghai/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	matchs := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		name := string(m[2])
		//result.Items 	= append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
							return ParseProfile(c, name)
						},
		})
		//fmt.Printf("User: %s,  URL: %s\n", m[2], m[1])
	}

	//获取其他页面   TO DO 去重复页面
	//matchs = cityUrlRe.FindAllSubmatch(contents, -1)
	//for _, m := range matchs {
	//	result.Requests = append(result.Requests, engine.Request{
	//		Url:		string(m[1]),
	//		ParserFunc: ParseCity,
	//	})
	//}
	return result
}
