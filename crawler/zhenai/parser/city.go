package parser

import (
	"regexp"
	"go-crawler/crawler/engine"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	// 将解析出来的 Url 列表都存储为一个ParserResult
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User " + string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: engine.NilParser,
			})
	}
	return result
}