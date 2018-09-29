package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)岁</td>`)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}


	match := ageRe.FindSubmatch(contents)
	age, err := strconv.Atoi(string(match[1]))
	if err != nil {
		profile.Age = age
	}

	match = marriageRe.FindSubmatch(contents)

	if err != nil {
		profile.Marriage = string(match[1])
	}
}

func extractString(contents byte[], re *Regexp){

}