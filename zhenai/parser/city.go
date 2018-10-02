package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe  = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		name := string(m[2])
		result.Items 	= append(result.Items, "User "+name)
		result.Requests = append(result.Requests, engine.Request{
			Url:		string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
							return ParseProfile(c, name)
						},
		})
		//fmt.Printf("User: %s,  URL: %s\n", m[2], m[1])
	}
	return result
}
