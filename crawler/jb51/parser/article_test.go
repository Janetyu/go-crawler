package parser

import (
	"testing"
	"go-crawler/crawler/fetcher"
	"go-crawler/crawler/utils"
)

func TestString(t *testing.T)  {
	//url := "https://www.jb51.net/article/73984.htm"
	//url2 := strings.Split(url, "/")
	//t.Logf("%d \n", len(url2))
	//t.Log(url2[len(url2)-1])
	//str := strings.Split(url2[len(url2)-1], ".")[0]
	//t.Log(str)

	s, err := fetcher.Fetch("https://www.jb51.net/article/145368.htm")
	if err != nil {
		panic(err)
	}

	t.Log("the title is ", utils.ExtractString([]byte(s), TitleRe))
	t.Log("the author is ", utils.ExtractString([]byte(s), AuthorRe))
	t.Log("the abstract is ", utils.ExtractString([]byte(s), AbstractRe))
	t.Log("the date is ", utils.ExtractString([]byte(s), DateRe))

}
