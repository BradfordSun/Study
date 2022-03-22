package parser

import (
	"crawler/engine"
	"regexp"
)

const articleListRe = `<a href="(/cn/premiumsupport/knowledge-center/[^"]+)"[^>]*>([^<]+)</a>`

func ParseUrlList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(articleListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), Parserfunc: engine.NilParser})
	}
	return result
}
