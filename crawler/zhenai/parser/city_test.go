// 测试city.go
package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCity(t *testing.T)  {
	content, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCity(content)

	expectedUrls := []string{
		"http://album.zhenai.com/u/1094206362",
		"http://album.zhenai.com/u/1769528526",
		"http://album.zhenai.com/u/1028744854",
	}
	//expectedUsers := []string{
	//	"User 若晨听海","User 夜明珠","User 梦的回忆",
	//}

	const resultSize  = 20
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",i, url, result.Requests[i].Url)
		}
	}

	//if len(result.Items) != resultSize {
	//	t.Errorf("result should have %d Items; but had %d",
	//		resultSize, len(result.Items))
	//}
	//
	//for i, item := range expectedUsers {
	//	if result.Items[i].(string) != item {
	//		t.Errorf("expected item #%d: %s; but was %s",i, item, result.Items[i].(string))
	//	}
	//}
}