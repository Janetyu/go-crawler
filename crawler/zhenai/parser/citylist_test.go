package parser

import (
	"testing"
	"go-crawler/crawler/fetcher"
	"fmt"
)

func TestParseCityList(t *testing.T)  {
	content, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", content)
}
