package parser

import (
	"regexp"
	"go-crawler/crawler/types"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]+>([^<]+)</a>`

func ParseCityList(contents []byte) types.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	
	// 将解析出来的 Url 列表都存储为一个ParserResult
	result := types.ParseResult{}
	//limit := 10
	for _, m := range matches {
		// m[0] 匹配的字符串本身
		//result.Items = append(result.Items, "City " + string(m[2]))
		result.Requests = append(result.Requests,
			types.Request{
				Url: string(m[1]),
				ParserFunc: ParseCity,
			})
		//limit--
		//if limit == 0 {
		//	break
		//}
		//fmt.Printf("City: %s, URL: %s \n",m[2], m[1])
	}
	//fmt.Printf("matches found : %d\n", len(matches))
	return result
}
