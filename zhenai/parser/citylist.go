package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		city := string(m[2])
		result.Items 	= append(result.Items, "City "+city)
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(m[1]),
			ParserFunc: ParseCity,
		})
		//fmt.Printf("City: %s,  URL: %s\n", m[2], m[1])
	}
	return result
}