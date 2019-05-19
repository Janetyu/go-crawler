package parser

import (
	"regexp"
	"go-crawler/crawler/types"
	"go-crawler/crawler/jb51/model"
	"go-crawler/crawler/utils"
	"strings"
)

var (
	TitleRe = regexp.MustCompile(`"title": "([^"]+)",`)
	DateRe = regexp.MustCompile(`更新时间：([^\s]+)`)
	AuthorRe = regexp.MustCompile(`[投稿|作者]+：([^\s]+)`)
	AbstractRe = regexp.MustCompile(`"description": "([^"]+)",`)
	LabelRe = regexp.MustCompile(`<li class="tag item"><[^>]+>([^<]+)</a></li>`)
	GolangType = "golang"
	PythonType = "python"
)

func ParseActicle(contents []byte, url string) types.ParseResult{
	// 将文章url中唯一的数字作为id
	str := strings.Split(url, "/")
	id := strings.Split(str[len(str)-1], ".")[0]

	matches := LabelRe.FindAllSubmatch(contents, -1)

	var label string
	for _, m := range matches {
		label = label + " " + string(m[1])
	}

	acticle := model.Article{
		Title: utils.ExtractString(contents, TitleRe),
		Author: utils.ExtractString(contents, AuthorRe),
		Abstract: utils.ExtractString(contents, AbstractRe),
		Label: label,
		Type: PythonType,
		Date: utils.ExtractString(contents, DateRe),
	}

	result := types.ParseResult{
		Items: []types.Item{
			{
				Url: url,
				Type: PythonType,
				Id: id,
				Payload: acticle,
			},
		},
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

