// 测试cityList.go
package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T)  {
	//content, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(content)

	const resultSize  = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝","阿克苏","阿拉善盟",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d Items; but had %d",
			resultSize, len(result.Items))
	}

	for i, item := range expectedCities {
		if result.Items[i].(string) != item {
			t.Errorf("expected item #%d: %s; but was %s",i, item, result.Items[i].(string))
		}
	}
}
