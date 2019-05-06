package persist

import (
	"testing"
	"go-crawler/crawler/zhenai/model"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"encoding/json"
)

func TestSave(t *testing.T) {
	excepted := model.Profile{
		Age: 29,
		MemberID: 1106374945,
		BasicInfo: `["未婚", "29岁", "天秤座(09.23-10.22)", "183cm", "95kg", "工作地:阿勒泰阿勒泰市", "月收入:8千-1.2万", "银行", "大学本科"]`,
		DetailInfo: `["汉族", "籍贯:山东德州", "体型:体格魁梧", "不吸烟", "社交场合会喝酒", "已购房", "已买车", "没有小孩", "是否想要孩子:想要孩子", "何时结婚:一年内"]`,
		Nickname: "寻真爱",
		GenderString: "男士",
		EducationString: "大学本科",
	}

	id, err := save(excepted)

	if err != nil {
		panic(err)
	}

	// must start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	//t.Logf("%+v", resp)
	t.Logf("%s", resp.Source)

	var actual model.Profile
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	if actual != excepted {
		t.Errorf("got %v; excepted %v",
			actual, excepted)
	}
}
