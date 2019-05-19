package parser

import (
	"go-crawler/crawler/types"
	"regexp"
	"strconv"
)

const (
	golangUrl = "/list/list_246_"
	pythonUrl = "/list/list_97_"
)

var (
	pageCount = 1
	ActicleRe = regexp.MustCompile(`<a href="(/article/[0-9]+.htm)"[^>]*>[^<]+</a>`)
	PageGolangRe = regexp.MustCompile(`<a href=/list/list_246_([0-9]+)[^>]+>末页</a>`)
	PagePythonRe = regexp.MustCompile(`<a href=/list/list_97_([0-9]+)[^>]+>末页</a>`)
	WebUrl = "https://www.jb51.net"
)

func ParseActicleList(contents []byte, url string) types.ParseResult {
	subMatches := PagePythonRe.FindSubmatch(contents)

	result := types.ParseResult{}

	// 表示无匹配
	if len(subMatches) < 2 {
		return result
	}

	// 总页数
	totalPage := string(subMatches[1])
	totalPageInt,err := strconv.Atoi(totalPage)
	if err != nil {
		return result
	}

	for ;pageCount <= totalPageInt; pageCount++ {
		url := WebUrl + pythonUrl + strconv.Itoa(pageCount) + ".htm"
		result.Requests = append(result.Requests,
			types.Request{
				Url: url,
				ParserFunc: ParseActicleList,
			})
	}

	allMatches := ActicleRe.FindAllSubmatch(contents, -1)
	for _,m := range allMatches {
		url := WebUrl + string(m[1])
		result.Requests = append(result.Requests,
			types.Request{
				Url: url,
				ParserFunc: ParseActicle,
			})
	}


	return result
}
