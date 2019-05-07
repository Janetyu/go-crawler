package parser

import (
	"regexp"
	"go-crawler/crawler/types"
)

var (
	profileRe  = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
			`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) types.ParseResult {
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents, -1)

	// 将解析出来的 Url 列表都存储为一个ParserResult
	result := types.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "User " + string(m[2]))
		url := string(m[1])
		result.Requests = append(result.Requests,
			types.Request{
				Url: url,
				ParserFunc: func(c []byte) types.ParseResult {
					return ParseProfile(c, url)
				},
			})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			types.Request{
				Url: string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}