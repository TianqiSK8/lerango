package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
)

const ageRe = `<td><span class="label">年龄：</span>([\d]+)岁</td>`

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)
	age, err := strconv.Atoi(string(match[1]))
	if err != nil {

	}
}
