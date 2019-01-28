package parser

import (
	"go-crawler/crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	// 将解析出来的 Url 列表都存储为一个ParserResult
	result := engine.ParseResult{}
	for _, m := range matches {
		// m[0] 匹配的字符串本身
		result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: ParseCity,
			})
		//fmt.Printf("City: %s, URL: %s \n",m[2], m[1])
	}
	//fmt.Printf("matches found : %d\n", len(matches))
	return result
}
