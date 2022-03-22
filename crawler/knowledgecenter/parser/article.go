package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
)

var updatedAtRe = regexp.MustCompile(`<i>上次更新日期：(.*)</i>`)
var titleRe = regexp.MustCompile(`<h1.*>(.*)</h1>`)

func parseArticle(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	profile.UpdatedAt = extractString(contents, updatedAtRe)
	profile.Title = extractString(contents, titleRe)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	if len(matches) >= 2 {
		return string(matches[1])
	} else {
		return ""
	}
}
