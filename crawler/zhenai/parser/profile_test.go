package parser

import (
	"testing"
	"go-crawler/crawler/zhenai/model"
	"go-crawler/crawler/fetcher"
	"go-crawler/crawler/types"
)

func TestParseProfile(t *testing.T) {
	content,_ := fetcher.Fetch("http://album.zhenai.com/u/1106374945")
	//e,_,_ := charset.DetermineEncoding(content,"")
	//t.Logf("the body encoding is %v\n",e)
	//t.Logf("html is : %s", content)
	//content, _ := ioutil.ReadFile("profile_test_data.html")

	//if err != nil {
	//	panic(err)
	//}

	result := ParseProfile(content, "http://album.zhenai.com/u/1106374945")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v",result.Items)
	}

	actual := result.Items[0]
	t.Log(actual)

	expected := types.Item{
		Url: "http://album.zhenai.com/u/1106374945",
		Type: "zhenai",
		Id: "1106374945",
		Payload: model.Profile{
			Age: 29,
			MemberID: 1106374945,
			BasicInfo: `["未婚", "29岁", "天秤座(09.23-10.22)", "183cm", "95kg", "工作地:阿勒泰阿勒泰市", "月收入:8千-1.2万", "银行", "大学本科"]`,
			DetailInfo: `["汉族", "籍贯:山东德州", "体型:体格魁梧", "不吸烟", "社交场合会喝酒", "已购房", "已买车", "没有小孩", "是否想要孩子:想要孩子", "何时结婚:一年内"]`,
			Nickname: "寻真爱",
			GenderString: "男士",
			EducationString: "大学本科",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
